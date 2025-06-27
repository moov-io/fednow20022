package common

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/errors"
)

// GetElement retrieves a field value from an item using a dot-notation path.
// Returns ErrFieldNotFound if the field doesn't exist.
// Returns ErrIndexOutOfBounds if array index is invalid.
func GetElement(item any, path string) (reflect.Type, any, error) {
	if item == nil {
		return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("item is nil: %w", errors.ErrFieldNotFound))
	}
	if path == "" {
		return nil, nil, errors.NewFieldError("", "get", fmt.Errorf("path is empty: %w", errors.ErrFieldNotFound))
	}

	v := reflect.ValueOf(item)

	// Dereference pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Make sure we're starting from a struct
	if v.Kind() != reflect.Struct {
		return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("item is not a struct: %w", errors.ErrFieldNotFound))
	}

	// Walk the path
	segments := strings.Split(path, ".")
	for _, segment := range segments {
		// Check if the segment is an array or slice access
		// e.g., "RptgReq[0]"
		re := regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
		matches := re.FindStringSubmatch(segment)
		if matches != nil {
			fieldName := matches[1]
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("invalid array index %q: %w", matches[2], errors.ErrIndexOutOfBounds))
			}
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			// Get the field by name
			if isReflectValueNil(v) {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("field %s is nil: %w", fieldName, errors.ErrFieldNotFound))
			}
			v = v.FieldByName(fieldName)
			if !v.IsValid() {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("field %s not found: %w", fieldName, errors.ErrFieldNotFound))
			}
			if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("field %s is not an array or slice: %w", fieldName, errors.ErrFieldNotFound))
			}

			// Check if the index is within bounds
			if index < 0 || index >= v.Len() {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("index %d out of bounds for field %s (length %d): %w", index, fieldName, v.Len(), errors.ErrIndexOutOfBounds))
			}

			// Access the element at the specified index
			v = v.Index(index)
		} else {
			// Regular field access
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() != reflect.Struct {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("cannot access field %s on non-struct type: %w", segment, errors.ErrFieldNotFound))
			}
			if isReflectValueNil(v) {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("parent field is nil: %w", errors.ErrFieldNotFound))
			}
			v = v.FieldByName(segment)
			if !v.IsValid() {
				return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("field %s not found: %w", segment, errors.ErrFieldNotFound))
			}
		}
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if isReflectValueNil(v) {
		return nil, nil, errors.NewFieldError(path, "get", fmt.Errorf("field value is nil: %w", errors.ErrFieldNotFound))
	}
	return v.Type(), v.Interface(), nil
}
func SetElementToModel(item any, path string, value any) error {
	return SetElementToDocument(item, path, value)
}
func SetElementToDocument(item any, path string, value any) error {
	if item == nil || path == "" {
		return fmt.Errorf("invalid input")
	}
	v := reflect.ValueOf(item)
	// item must be a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("item must be a pointer to a struct")
	}
	// Dereference to struct
	v = v.Elem()
	segments := strings.Split(path, ".")
	for i := 0; i < len(segments)-1; i++ {
		re := regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
		matches := re.FindStringSubmatch(segments[i])
		if matches != nil {
			fieldName := matches[1]
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil // Invalid index
			}
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if isReflectValueNil(v) {
				return nil // Field is nil
			}
			v = v.FieldByName(fieldName)
			if !v.IsValid() || (v.Kind() != reflect.Slice && v.Kind() != reflect.Array) {
				return nil // Field not found or not a slice/array
			}
			if isEmpty(v) {
				newArray := reflect.New(v.Type()).Elem()
				v.Set(newArray)
			}
			if index >= v.Len() {
				elementType := v.Type().Elem()

				// Handle pointer types
				if elementType.Kind() == reflect.Ptr {
					elementType = elementType.Elem()
				}

				// Ensure the underlying type is a struct
				if elementType.Kind() == reflect.Struct {
					// Create a new slice with the required length
					newSlice := reflect.MakeSlice(v.Type(), index+1, index+1)
					reflect.Copy(newSlice, v)

					// Initialize new elements in the slice
					for i := v.Len(); i <= index; i++ {
						newStruct := reflect.New(elementType).Elem()
						if v.Type().Elem().Kind() == reflect.Ptr {
							// If the slice holds pointers, set a pointer to the new struct
							newSlice.Index(i).Set(newStruct.Addr())
						} else {
							// Otherwise, set the struct directly
							newSlice.Index(i).Set(newStruct)
						}
					}

					// Replace the old slice with the new slice
					v.Set(newSlice)
				} else {
					return fmt.Errorf("element type is not a struct or pointer to a struct")
				}
			}
			if index < v.Len() {
				v = v.Index(index)
			}
		} else {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if isReflectValueNil(v) {
				return fmt.Errorf("field %s is nil", segments[i])
			}
			field := v.FieldByName(segments[i])
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", segments[i])
			}
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", segments[i])
			}
			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				field = field.Elem()
			}
			// Move deeper
			if field.Kind() != reflect.Struct {
				return fmt.Errorf("field %s is not a struct", segments[i])
			}

			v = field
		}
	}
	// Now set the last field
	last := segments[len(segments)-1]
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if isReflectValueNil(v) {
		return fmt.Errorf("field %s is nil", last)
	}

	/*set value to last type is array*/
	re := regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
	matches := re.FindStringSubmatch(last)
	if matches != nil {
		fieldName := matches[1]
		index, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil // Invalid index
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if isReflectValueNil(v) {
			return nil // Field is nil
		}
		v = v.FieldByName(fieldName)
		if !v.IsValid() || (v.Kind() != reflect.Slice && v.Kind() != reflect.Array) {
			return nil // Field not found or not a slice/array
		}
		if isEmpty(v) {
			newArray := reflect.New(v.Type()).Elem()
			v.Set(newArray)
		}
		if index >= v.Len() {
			elementType := v.Type().Elem()

			// Handle pointer types
			if elementType.Kind() == reflect.Ptr {
				elementType = elementType.Elem()
			}

			newSlice := reflect.MakeSlice(v.Type(), index+1, index+1)
			reflect.Copy(newSlice, v)

			// Initialize new elements in the slice
			for i := v.Len(); i <= index; i++ {
				newStruct := reflect.New(elementType).Elem()
				if v.Type().Elem().Kind() == reflect.Ptr {
					// If the slice holds pointers, set a pointer to the new struct
					newSlice.Index(i).Set(newStruct.Addr())
				} else {
					// Otherwise, set the struct directly
					newSlice.Index(i).Set(newStruct)
				}
			}

			// Replace the old slice with the new slice
			v.Set(newSlice)
		}
		if index < v.Len() {
			field := v.Index(index)
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", last)
			}
			if !field.CanSet() {
				return fmt.Errorf("field %s cannot be set (may be unexported)", last)
			}

			// Convert value if necessary
			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				field = field.Elem()
			}
			err := setValue(field, value)
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		field := v.FieldByName(last)
		if !field.IsValid() {
			return fmt.Errorf("field %s not found", last)
		}
		if !field.CanSet() {
			return fmt.Errorf("field %s cannot be set (may be unexported)", last)
		}

		// Convert value if necessary
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			field = field.Elem()
		}
		err := setValue(field, value)
		if err != nil {
			return err
		}
	}
	return nil
}
func setValue(v reflect.Value, value any) error {
	if !v.CanSet() {
		return fmt.Errorf("cannot set value")
	}
	val := reflect.ValueOf(value)
	if val.Type().ConvertibleTo(v.Type()) {
		v.Set(val.Convert(v.Type()))
		if hasValidateMethod(v) {
			method := v.MethodByName("Validate")
			if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
				// Call the Validate method
				results := method.Call(nil) //nolint:forbidigo
				if len(results) == 1 && !results[0].IsNil() {
					validationErr, ok := results[0].Interface().(error)
					if ok {
						return validationErr
					}
					return fmt.Errorf("%v", results[0].Interface()) // Fallback for non-error types
				}
			}
		}
	} else if val.Type().Kind() == reflect.String && v.Type().Kind() == reflect.String {
		if strVal, ok := val.Interface().(string); ok {
			convertedVal := reflect.ValueOf(strVal).Convert(v.Type())
			v.Set(convertedVal)
			if hasValidateMethod(v) {
				method := v.MethodByName("Validate")
				if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
					// Call the Validate method
					results := method.Call(nil) //nolint:forbidigo
					if len(results) == 1 && !results[0].IsNil() {
						validationErr, ok := results[0].Interface().(error)
						if ok {
							return validationErr
						}
						return fmt.Errorf("%v", results[0].Interface()) // Fallback for non-error types
					}
				}
			}
		} else {
			return fmt.Errorf("value is not a string, cannot convert to field type %s", v.Type())
		}
	} else if val.Type() == reflect.TypeOf(fednow.ISODate{}) && v.Type().Kind() == reflect.String {
		/*Convert pacs_008_001_10.ISOYear to fednow.ISODate*/
		isoDate, ok := val.Interface().(fednow.ISODate)
		if !ok {
			return fmt.Errorf("failed to convert value to fednow.ISODate")
		}
		yearStr := fmt.Sprintf("%d", isoDate.Year)
		convertedVal := reflect.ValueOf(yearStr).Convert(v.Type())
		v.Set(convertedVal)
		if hasValidateMethod(v) {
			method := v.MethodByName("Validate")
			if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
				// Call the Validate method
				results := method.Call(nil) //nolint:forbidigo
				if len(results) == 1 && !results[0].IsNil() {
					validationErr, ok := results[0].Interface().(error)
					if ok {
						return validationErr
					}
					return fmt.Errorf("%v", results[0].Interface()) // Fallback for non-error types
				}
			}
		}
	} else if val.Kind() == reflect.String && v.Type() == reflect.TypeOf(fednow.ISODate{}) {
		// Convert string to fednow.ISODate
		strVal, ok := val.Interface().(string)
		if !ok {
			return fmt.Errorf("value is not a string, cannot convert to fednow.ISODate")
		}

		// Assuming fednow.ISODate.Year is an integer
		year, err := strconv.Atoi(strVal)
		if err != nil {
			return fmt.Errorf("failed to convert string to integer for ISODate.Year: %w", err)
		}

		var isoDate fednow.ISODate
		isoDate.Year = year // Assign the converted integer value

		v.Set(reflect.ValueOf(isoDate))
	} else if val.Kind() == reflect.String {
		// Handle string to numeric conversions
		strVal, ok := val.Interface().(string)
		if !ok {
			return fmt.Errorf("cannot convert value to field type %s", v.Type())
		}

		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				return fmt.Errorf("cannot convert string %q to %s: %w", strVal, v.Type(), err)
			}
			v.SetInt(intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(strVal, 10, 64)
			if err != nil {
				return fmt.Errorf("cannot convert string %q to %s: %w", strVal, v.Type(), err)
			}
			v.SetUint(uintVal)
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(strVal, 64)
			if err != nil {
				return fmt.Errorf("cannot convert string %q to %s: %w", strVal, v.Type(), err)
			}
			v.SetFloat(floatVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(strVal)
			if err != nil {
				return fmt.Errorf("cannot convert string %q to %s: %w", strVal, v.Type(), err)
			}
			v.SetBool(boolVal)
		case reflect.String:
			v.SetString(strVal)
		case reflect.Invalid, reflect.Uintptr, reflect.Complex64, reflect.Complex128,
			reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
			return fmt.Errorf("cannot convert string to field type %s", v.Type())
		}
	} else {
		return fmt.Errorf("cannot convert value to field type %s", v.Type())
	}

	return nil
}

func hasValidateMethod(v reflect.Value) bool {
	// Get the type of the value
	t := v.Type()

	// Check if the method "Validate" exists
	method, exists := t.MethodByName("Validate")
	if !exists {
		return false
	}

	// Ensure the method has the correct signature (e.g., no parameters and returns an error)
	if method.Type.NumIn() == 1 && method.Type.NumOut() == 1 && method.Type.Out(0) == reflect.TypeOf((*error)(nil)).Elem() {
		return true
	}

	return false
}
func RemakeMapping(from any, modelMap map[string]any, toModel bool) map[string]string {
	newMap := make(map[string]string)

	for k, v := range modelMap {
		switch typedVal := v.(type) {
		case string:
			newMap[k] = typedVal

		case map[string]string:
			processFlatSliceMapping(from, newMap, k, typedVal, toModel)

		case map[string]any:
			processNestedSliceMapping(from, newMap, k, typedVal, toModel)
		}
	}

	return newMap
}
func processFlatSliceMapping(from any, result map[string]string, key string, mapping map[string]string, toModel bool) {
	src, dst := seperateKeyAndValue(key, ":")
	if src == "" || dst == "" {
		return
	}
	targetPath := strings.TrimSpace(dst)
	if toModel {
		targetPath = strings.TrimSpace(src)
	}

	_, val, err := GetElement(from, targetPath)
	if err != nil {
		return // Silently ignore errors for backward compatibility
	}
	valValue := reflect.ValueOf(val)
	if val == nil || (valValue.Kind() != reflect.Array && valValue.Kind() != reflect.Slice) {
		return
	}

	for i := 0; i < valValue.Len(); i++ {
		for k1, v1 := range mapping {
			result[fmt.Sprintf("%s[%d].%s", src, i, k1)] = fmt.Sprintf("%s[%d].%s", dst, i, v1)
		}
	}
}

func processNestedSliceMapping(from any, result map[string]string, key string, mapping map[string]any, toModel bool) {
	src, dst := seperateKeyAndValue(key, ":")
	if src == "" || dst == "" {
		return
	}
	targetPath := strings.TrimSpace(dst)
	if toModel {
		targetPath = strings.TrimSpace(src)
	}

	_, val, err := GetElement(from, targetPath)
	if err != nil {
		return // Silently ignore errors for backward compatibility
	}
	valValue := reflect.ValueOf(val)
	if val == nil || (valValue.Kind() != reflect.Array && valValue.Kind() != reflect.Slice) {
		return
	}

	for i := 0; i < valValue.Len(); i++ {
		for k1, v1 := range mapping {
			switch inner := v1.(type) {
			case string:
				if k1 == "index" && inner == "index" { /// string array
					result[fmt.Sprintf("%s[%d]", src, i)] = fmt.Sprintf("%s[%d]", dst, i)
				} else {
					result[fmt.Sprintf("%s[%d].%s", src, i, k1)] = fmt.Sprintf("%s[%d].%s", dst, i, inner)
				}
			case map[string]string:
				src2, dst2 := seperateKeyAndValue(k1, ":")
				if src2 == "" || dst2 == "" {
					continue
				}
				targetPath2 := strings.TrimSpace(dst2)
				if toModel {
					targetPath2 = strings.TrimSpace(src2)
				}
				if toModel {
					targetPath2 = fmt.Sprintf("%s[%d].%s", src, i, targetPath2)
				} else {
					targetPath2 = fmt.Sprintf("%s[%d].%s", dst, i, targetPath2)
				}
				_, val2, err := GetElement(from, targetPath2)
				if err != nil {
					continue // Silently ignore errors for backward compatibility
				}
				valValue2 := reflect.ValueOf(val2)
				if val2 == nil || (valValue2.Kind() != reflect.Array && valValue2.Kind() != reflect.Slice) {
					continue
				}

				for j := 0; j < valValue2.Len(); j++ {
					for k2, v2 := range inner {
						result[fmt.Sprintf("%s[%d].%s[%d].%s", src, i, src2, j, k2)] =
							fmt.Sprintf("%s[%d].%s[%d].%s", dst, i, dst2, j, v2)
					}
				}
			}
		}
	}
}
func seperateKeyAndValue(src string, separate string) (string, string) {
	parts := strings.Split(src, separate)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	if len(parts) == 1 {
		return strings.TrimSpace(parts[0]), ""
	}
	return "", ""
}

func isReflectValueNil(v reflect.Value) bool {
	// First check if the reflect.Value is valid
	if !v.IsValid() {
		return true
	}

	// Then check if the kind supports nil and if it's actually nil
	switch v.Kind() {
	case reflect.Invalid, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128, reflect.Array, reflect.String, reflect.Struct,
		reflect.UnsafePointer:
		return false

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
func isEmpty(val interface{}) bool {
	switch v := val.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int() == 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint() == 0
	case float32, float64:
		return reflect.ValueOf(v).Float() == 0
	case bool:
		return !v
	case time.Time:
		return v.IsZero()
	default:
		// Use reflect for fallback
		return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
	}
}
