package fednow_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/fednow20022/gen/pacs_008_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow"

	"github.com/stretchr/testify/require"
)

func TestValidateMultipleErrors(t *testing.T) {
	var errs base.ErrorList = base.ErrorList{}
	fednow.AddError(&errs, "Max35Text", pacs_008_001_08.Max35Text("B20230931145322200000057A11712044729M").Validate())
	require.Len(t, errs, 1)
	require.ErrorContains(t, errs.Err(), "Max35Text: B20230931145322200000057A11712044729M fails validation with length 37 <= required maxLength 35")
}

func TestValidatePattern(t *testing.T) {
	pattern := `[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))((([0-1][0-9])|(2[0-3]))(([0-5][0-9])){2})[A-Z0-9]{11}.*`
	require.NoError(t, fednow.ValidatePattern("20230713145322200000057A11712044729", pattern))
	require.Error(t, fednow.ValidatePattern("20230931145322200000057A11712044729", pattern)) // invalid MMDD

	t.Run("concurrent", func(t *testing.T) {
		iterations := 1000

		var wg sync.WaitGroup
		wg.Add(iterations)

		for i := 0; i < iterations; i++ {
			go func(idx int) {
				wg.Done()

				patern := fmt.Sprintf("[0-9]{0,%d}", idx/5) // add/get regexes from cache
				require.NoError(t, fednow.ValidatePattern("4", patern))
			}(i)
		}

		wg.Wait()
	})
}

func TestValidateEnumeration(t *testing.T) {
	enumVals := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	require.NoError(t, fednow.ValidateEnumeration("A", enumVals...))
	require.Error(t, fednow.ValidateEnumeration("Z", enumVals...))
}

func TestValidateLength(t *testing.T) {
	length := 5
	require.NoError(t, fednow.ValidateLength("abcde", length))
	require.Error(t, fednow.ValidateLength("abcdef", length))
	require.NoError(t, fednow.ValidateLength("abµde", length))
	require.Error(t, fednow.ValidateLength("abµdef", length))
}

func TestValidateMinLength(t *testing.T) {
	minLength := 5
	require.NoError(t, fednow.ValidateMinLength("abcde", minLength))
	require.Error(t, fednow.ValidateMinLength("abcd", minLength))
	require.NoError(t, fednow.ValidateMinLength("abµde", minLength))
	require.Error(t, fednow.ValidateMinLength("abµd", minLength))
}

func TestValidateMaxLength(t *testing.T) {
	maxLength := 5
	require.NoError(t, fednow.ValidateMaxLength("abcde", maxLength))
	require.Error(t, fednow.ValidateMaxLength("abcdef", maxLength))
	require.NoError(t, fednow.ValidateMaxLength("abµde", maxLength))
	require.Error(t, fednow.ValidateMaxLength("abµdef", maxLength))
}

func TestValidateMinInclusive(t *testing.T) {
	minVal := 3.0
	require.Error(t, fednow.ValidateMinInclusive(3.0, minVal))
	require.NoError(t, fednow.ValidateMinInclusive(5.0, minVal))
	require.Error(t, fednow.ValidateMinInclusive(2.0, minVal))
}

func TestValidateMaxInclusive(t *testing.T) {
	maxVal := 3
	require.Error(t, fednow.ValidateMaxInclusive(3, maxVal))
	require.NoError(t, fednow.ValidateMaxInclusive(2, maxVal))
	require.Error(t, fednow.ValidateMaxInclusive(5, maxVal))
}

func TestValidateMinExclusive(t *testing.T) {
	minVal := 3
	require.NoError(t, fednow.ValidateMinExclusive(3, minVal))
	require.NoError(t, fednow.ValidateMinExclusive(5, minVal))
	require.Error(t, fednow.ValidateMinExclusive(2, minVal))
}

func TestValidateMaxExclusive(t *testing.T) {
	maxVal := 3
	require.NoError(t, fednow.ValidateMaxExclusive(3, maxVal))
	require.NoError(t, fednow.ValidateMaxExclusive(2, maxVal))
	require.Error(t, fednow.ValidateMaxExclusive(5, maxVal))
}

func TestValidateFractionDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, fednow.ValidateFractionDigits("3.1415", maxVal))
	require.NoError(t, fednow.ValidateFractionDigits("3.14159", maxVal))
	require.Error(t, fednow.ValidateFractionDigits("3.141592", maxVal))
}

func TestValidateTotalDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, fednow.ValidateTotalDigits("3.141", maxVal))
	require.NoError(t, fednow.ValidateTotalDigits("3.1415", maxVal))
	require.Error(t, fednow.ValidateTotalDigits("3.14159", maxVal))
}
