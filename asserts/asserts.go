package asserts

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func Slices[T cmp.Ordered](t *testing.T, expectedSlice []T, resultSlice []T) {
	slices.Sort(expectedSlice)
	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf("expected %v, got %v", expectedSlice, resultSlice)
	}
}

func Equal[T comparable](t *testing.T, expected T, result T) {
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
