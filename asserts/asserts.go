package asserts

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

type Tester interface {
	Errorf(format string, args ...interface{})
}

func Slices[T cmp.Ordered](t *testing.T, expectedSlice []T, resultSlice []T) {
	slices.Sort(expectedSlice)
	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf("expected %v, got %v", expectedSlice, resultSlice)
	}
}

func DeepEqual[T any](t *testing.T, expected, result T) {
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Equal[T comparable](t Tester, expected T, result T) {
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
