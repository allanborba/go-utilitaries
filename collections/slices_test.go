package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
	"github.com/allanborba/utilitaries/collections"
)

func TestSlicesContains(t *testing.T) {
	type TestStruct struct {
		A int
		b string
	}

	sliceOfInt := []int{1, 2, 3}
	sliceOfStruct := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
		{A: 3, b: "c"},
	}

	asserts.Equal(t, true, collections.Contains(sliceOfInt, 1))
	asserts.Equal(t, true, collections.Contains(sliceOfInt, 2))
	asserts.Equal(t, true, collections.Contains(sliceOfInt, 3))
	asserts.Equal(t, false, collections.Contains(sliceOfInt, 4))

	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 1, b: "a"}))
	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 2, b: "b"}))
	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 3, b: "c"}))
	asserts.Equal(t, false, collections.Contains(sliceOfStruct, TestStruct{A: 4, b: "d"}))
}
