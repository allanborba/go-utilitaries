package asserts

import (
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"testing"
)

const ERROR_MSG = "expected %v, got %v"

type Tester interface {
	Errorf(format string, args ...interface{})
}

func Equal[T comparable](t Tester, expected T, result T) {
	if result != expected {
		t.Errorf(ERROR_MSG, expected, result)
	}
}

func Slices[T cmp.Ordered](t *testing.T, expectedSlice []T, resultSlice []T) {
	slices.Sort(expectedSlice)
	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf(ERROR_MSG, expectedSlice, resultSlice)
	}
}

func DeepEqual[T comparable](t Tester, expected, result T) {
	if isStruct(expected) {
		compareStructs(expected, result, t)
	} else {
		Equal(t, expected, result)
	}
}

func compareStructs[T any](expected T, result T, t Tester) {
	fields := GetFieldNames(expected)
	expectedMap := StructToMap(expected)
	resultMap := StructToMap(result)

	errorsKeys := []string{}
	expectedMsg := ""
	resultMsg := ""

	for _, field := range fields {
		if expectedMap[field] == resultMap[field] {
			continue
		}

		expectedMsg += fmt.Sprintf("%v: %v ", field, expectedMap[field])
		resultMsg += fmt.Sprintf("%v: %v ", field, resultMap[field])
		errorsKeys = append(errorsKeys, field)
	}

	if len(errorsKeys) > 0 {
		expectedMsg = fmt.Sprint("expected { ", expectedMsg, "}")
		resultMsg = fmt.Sprint("got { ", resultMsg, "}")

		t.Errorf(fmt.Sprintf("%v, %v", expectedMsg, resultMsg))
	}
}

func isStruct[T any](expected T) bool {
	value := reflect.ValueOf(expected)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value.Kind() == reflect.Struct
}
