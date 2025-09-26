package asserts

import (
	"fmt"
)

func DeepEqual[T comparable](t Tester, expected, result T) {
	if IsStruct(expected) {
		assertStructs(expected, result, t)
	} else {
		Equal(t, expected, result)
	}
}

func assertStructs[T any](expected T, result T, t Tester) {
	expectedMsg, resultMsg := compareStructs(expected, result)

	if expectedMsg != "" {
		expectedMsg = fmt.Sprint("expected ", expectedMsg)
		resultMsg = fmt.Sprint("got ", resultMsg)

		t.Errorf(fmt.Sprintf("%v, %v", expectedMsg, resultMsg))
	}
}

func compareStructs[T any](expected T, result T) (string, string) {
	fields := GetFieldNames(expected)
	expectedMap := StructToMap(expected)
	resultMap := StructToMap(result)

	expectedMsg := ""
	resultMsg := ""

	for _, field := range fields {
		if IsStruct(expectedMap[field]) {
			innerExpectedMsg, innerResultMsg := compareStructs(expectedMap[field], resultMap[field])

			if innerExpectedMsg != "" {
				expectedMsg += fmt.Sprintf("%v: %v ", field, innerExpectedMsg)
				resultMsg += fmt.Sprintf("%v: %v ", field, innerResultMsg)
			}

			continue
		}

		if expectedMap[field] != resultMap[field] {
			expectedMsg += fmt.Sprintf("%v: %v ", field, expectedMap[field])
			resultMsg += fmt.Sprintf("%v: %v ", field, resultMap[field])
		}
	}

	if expectedMsg != "" {
		expectedMsg = fmt.Sprint("{ ", expectedMsg, "}")
		resultMsg = fmt.Sprint("{ ", resultMsg, "}")
	}

	return expectedMsg, resultMsg
}
