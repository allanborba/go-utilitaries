package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/collections"
)

func TestIinitializeSetWithCorrectlyLength(t *testing.T) {
	assertSetLength(t, 0, collections.NewSet([]int{}).Len())
	assertSetLength(t, 3, collections.NewSet([]int{1, 2, 3}).Len())
	assertSetLength(t, 2, collections.NewSet([]string{"1", "2"}).Len())
}

func TestAddElementOnSet(t *testing.T) {
	set := collections.NewSet([]int{})
	set.Add(5)
	set.Add(10)

	assertSetLength(t, 2, set.Len())
}

func TestWhenTryToAddRepitedElementDoesNotDuplicatedOnSet(t *testing.T) {
	set := collections.NewSet([]int{})
	set.Add(5)
	set.Add(10)
	set.Add(10)
	set.Add(5)

	assertSetLength(t, 2, set.Len())
}

func assertSetLength(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
