package asserts

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/allanborba/go-utilitaries/collections"
)

func Object[T any](t Tester, expected, result T) {
	if IsStruct(expected) {
		ObjectIgnoringFields(t, expected, result, []string{})
	} else {
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}
}

func ObjectIgnoringFields[T any](t Tester, expected T, result T, fieldsToIgnore []string) {
	hasError, expectedMsg, resultMsg := compareStruct(expected, result, fieldsToIgnore)
	if hasError {
		t.Errorf("expected %s, got %s", expectedMsg, resultMsg)
	}
}

func compareStruct[T any](expected T, result T, fieldsToIgnore []string) (bool, string, string) {
	var hasError bool

	wrongValuesExpected := make([]string, 0)
	wrongValuesResult := make([]string, 0)
	mapExpected := StructToMap(expected)
	mapResult := StructToMap(result)
	fields := GetFieldNames(expected)

	for _, field := range fields {
		var expectedMsg, resultMsg string
		if collections.Contains(fieldsToIgnore, field) {
			continue
		}

		if IsStruct(mapExpected[field]) {
			hasError, expectedMsg, resultMsg = compareStruct(mapExpected[field], mapResult[field], fieldsToIgnore)
			expectedMsg = fmt.Sprintf("%s: %s", field, expectedMsg)
			resultMsg = fmt.Sprintf("%s: %s", field, resultMsg)
			wrongValuesExpected = append(wrongValuesExpected, expectedMsg)
			wrongValuesResult = append(wrongValuesResult, resultMsg)
		} else {
			if !reflect.DeepEqual(mapExpected[field], mapResult[field]) {
				hasError = true
				expectedMsg = fmt.Sprintf("%s: %v", field, mapExpected[field])
				resultMsg = fmt.Sprintf("%s: %v", field, mapResult[field])
				wrongValuesExpected = append(wrongValuesExpected, expectedMsg)
				wrongValuesResult = append(wrongValuesResult, resultMsg)
			}
		}
	}

	if !hasError {
		return false, "", ""
	}

	expectedMsg := fmt.Sprintf("{%s}", strings.Join(wrongValuesExpected, ", "))
	resultMsg := fmt.Sprintf("{%s}", strings.Join(wrongValuesResult, ", "))

	return hasError, expectedMsg, resultMsg
}
