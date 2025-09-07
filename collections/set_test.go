package collections_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/allanborba/utilitaries/collections"
)

func TestIinitializeSetWithCorrectlyLength(t *testing.T) {
	assertSetLength(t, 0, collections.NewSet([]int{}).Len())
	assertSetLength(t, 3, collections.NewSet([]int{1, 2, 3}).Len())
	assertSetLength(t, 2, collections.NewSet([]string{"1", "2"}).Len())
}

func TestAddElementsOnlyOnceOnSet(t *testing.T) {
	set := collections.NewSet([]int{})
	set.Add(5)
	set.Add(10)
	set.Add(10)
	set.Add(5)

	assertSetLength(t, 2, set.Len())
}

func TestConvertSetInToSlice(t *testing.T) {
	set := collections.NewSet([]int{1, 2, 3, 10, 15})

	expectedSlice := []int{1, 2, 3, 10, 15}
	resultSlice := set.ToSlice()

	slices.Sort(resultSlice)

	if !reflect.DeepEqual(expectedSlice, resultSlice) {
		t.Errorf("expected %d, got %d", expectedSlice, resultSlice)
	}
}

func assertSetLength(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
