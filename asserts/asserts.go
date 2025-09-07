package asserts

import (
	"reflect"
	"slices"
	"testing"
)

func Slices(t *testing.T, expectedSlice []int, resultSlice []int) {
	slices.Sort(expectedSlice)
	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf("expected %d, got %d", expectedSlice, resultSlice)
	}
}

func Equal[T comparable](t *testing.T, expected T, result T) {
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
