package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
	"github.com/allanborba/utilitaries/collections"
)

func TestIinitializeSetWithCorrectlyLength(t *testing.T) {
	asserts.Equal(t, 0, collections.NewSet([]int{}).Len())
	asserts.Equal(t, 3, collections.NewSet([]int{1, 2, 3}).Len())
	asserts.Equal(t, 2, collections.NewSet([]string{"1", "2"}).Len())
}

func TestAddElementsOnlyOnceOnSet(t *testing.T) {
	set := collections.NewSet([]int{})
	set.Add(5)
	set.Add(10)
	set.Add(10)
	set.Add(5)

	asserts.Equal(t, 2, set.Len())
}

func TestConvertSetInToSlice(t *testing.T) {
	resultSlice := collections.NewSet([]int{1, 2, 3, 10, 15}).ToSlice()
	expectedSlice := []int{1, 2, 3, 10, 15}

	asserts.Slices(t, expectedSlice, resultSlice)
}

func TestInformCorrectlyIfHasElementOnSet(t *testing.T) {
	set := collections.NewSet([]int{1, 2, 3})

	asserts.Equal(t, true, set.Has(1))
}
