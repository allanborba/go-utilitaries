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

func Slices[T cmp.Ordered](t *testing.T, expectedSlice []T, resultSlice []T) {
	slices.Sort(expectedSlice)
	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf(ERROR_MSG, expectedSlice, resultSlice)
	}
}

func DeepEqual[T any](t Tester, expected, result T) {
	if !reflect.DeepEqual(expected, result) {
		t.Errorf(ERROR_MSG, expected, result)
	}
}

func Equal[T comparable](t Tester, expected T, result T) {
	if result != expected {
		t.Errorf(ERROR_MSG, expected, result)
	}
}

func GetFieldNames[T any](obj T) []string {
	fields := []string{}
	value := reflect.ValueOf(obj)
	typeOf := value.Type()

	for i := range typeOf.NumField() {
		filedName := typeOf.Field(i).Name
		fields = append(fields, filedName)
	}

	return fields
}
