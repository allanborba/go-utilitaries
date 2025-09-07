package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/collections"
)

func TestIinitializeEmptySet(t *testing.T) {
	set := collections.NewSet([]int{})

	if set.Len() != 0 {
		t.Errorf("expected 0, got %d", set.Len())
	}
}

func TestIinitializeSetOfIntegers(t *testing.T) {
	set := collections.NewSet([]int{1, 2, 3})

	if set.Len() != 3 {
		t.Errorf("expected 0, got %d", set.Len())
	}
}

func TestInitializeSetOfStrings(t *testing.T) {
	set := collections.NewSet([]string{"1", "2", "3", "4"})

	if set.Len() != 4 {
		t.Errorf("expected 0, got %d", set.Len())
	}
}
