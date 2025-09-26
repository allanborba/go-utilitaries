package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
	"github.com/allanborba/utilitaries/collections"
)

func TestSlicesContains(t *testing.T) {
	slice1 := []int{1, 2, 3}

	asserts.Equal(t, true, collections.Contains(slice1, 1))
	asserts.Equal(t, true, collections.Contains(slice1, 2))
	asserts.Equal(t, true, collections.Contains(slice1, 3))
	asserts.Equal(t, false, collections.Contains(slice1, 4))
}
