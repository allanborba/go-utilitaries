package asserts

import (
	"cmp"
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
		Equal(t, expected, result)
	} else {
		compareStructs(expected, result, t)
	}
}

func compareStructs[T any](expected T, result T, t Tester) {
	fields := GetFieldNames(expected)
	expectedMap := StructToMap(expected)
	resultMap := StructToMap(result)

	for _, field := range fields {
		if expectedMap[field] != resultMap[field] {
			t.Errorf("expected {%v: %v}, got {%v: %v}", field, expectedMap[field], field, resultMap[field])
		}
	}
}

func isStruct[T any](expected T) bool {
	value := reflect.ValueOf(expected)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	isStruct := value.Kind() != reflect.Struct
	return isStruct
}
