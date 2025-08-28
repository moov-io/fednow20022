package common

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/pkg/fednow/errors"
)

type Match struct {
	SrcPath string
	DstPath string
}
type ElementMap struct {
	SrcElement string
	DstElement string
}

type Document struct {
	Attrs []xml.Attr `xml:",any,attr"`
}
type ISODocument interface {
	// Validate() error
}
type DocumentFactory func() ISODocument
type DataFactory func() any

// DocumentFrom parses XML data and creates an ISODocument using the appropriate factory.
// Returns ErrInvalidXML if XML parsing fails.
// Returns ErrUnknownNamespace if the XML namespace is not recognized.
func DocumentFrom(data []byte, factory DocumentFactory) (ISODocument, string, error) {
	var root Document
	if err := xml.Unmarshal(data, &root); err != nil {
		return nil, "", errors.NewParseError("XML decode", "document", err)
	}

	var xmlns string
	for _, attr := range root.Attrs {
		if attr.Name.Local == "xmlns" && attr.Name.Space == "" {
			xmlns = attr.Value
			break
		}
	}

	if xmlns == "" {
		return nil, "", fmt.Errorf("XML document missing xmlns attribute: %w", errors.ErrInvalidXML)
	}

	// Instantiate and unmarshal into actual model
	doc := factory()
	if err := xml.Unmarshal(data, doc); err != nil {
		return nil, "", errors.NewParseError("XML unmarshal", "model structure", err)
	}

	return doc, xmlns, nil
}

func MessageModelWith(doc ISODocument, factory DataFactory, pathMap map[string]any) (interface{}, error) {
	if pathMap == nil {
		return nil, fmt.Errorf("pathMap cannot be nil")
	}
	dataModel := factory()
	rePathMap := RemakeMapping(doc, pathMap, true)
	for sourcePath, targetPath := range rePathMap {
		CopyDocumentValueToMessage(doc, sourcePath, dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model interface{}, pathMap map[string]any, factory DocumentFactory) (ISODocument, error) {
	if model == nil {
		return nil, fmt.Errorf("model cannot be nil")
	}
	document := factory()
	rePathMap := RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
		if err := CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, err
		}
	}
	return document, nil
}
func CheckRequireFields(model interface{}, requireFields []string) error {
	if model == nil {
		return fmt.Errorf("model cannot be nil")
	}
	for _, field := range requireFields {
		_, value, err := GetElement(model, field)
		if err != nil {
			return errors.NewRequiredFieldError(field)
		}
		if isEmpty(value) {
			return errors.NewRequiredFieldError(field)
		}
	}
	return nil
}
func XMLFrom(doc ISODocument) ([]byte, error) {
	if doc == nil {
		return nil, fmt.Errorf("document cannot be nil")
	}
	data, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document: %w", err)
	}
	return data, nil
}
func JSONFrom(model interface{}) ([]byte, error) {
	if model == nil {
		return nil, fmt.Errorf("model cannot be nil")
	}
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal model: %w", err)
	}
	return data, nil
}
func CopyDocumentValueToMessage(from any, fromPah string, to any, toPath string) {
	if from == nil || fromPah == "" || toPath == "" {
		return
	}
	_, value, err := GetElement(from, fromPah)
	if err != nil {
		return // Silently ignore field access errors for backward compatibility
	}
	if isEmpty(value) {
		return
	}
	if value == nil {
		return
	}

	err = SetElementToModel(to, toPath, value)
	if err != nil {
		return
	}
}

func CopyMessageValueToDocument(from any, fromPath string, to ISODocument, toPath string) error {
	if from == nil || fromPath == "" || toPath == "" {
		return fmt.Errorf("invalid input")
	}
	_, value, err := GetElement(from, fromPath)
	if err != nil {
		return fmt.Errorf("failed to get field %s: %w", fromPath, err)
	}
	if isEmpty(value) {
		return nil
	}

	err = SetElementToDocument(to, toPath, value)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", fromPath, err)
	}
	return nil
}
