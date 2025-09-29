package collections_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
	"github.com/allanborba/utilitaries/collections"
)

type TestStruct struct {
	A int
	b string
}

func TestSlicesContains(t *testing.T) {
	sliceOfInt := buildSliceOfIntsHelper()
	sliceOfStruct := buildSliceOfStructsHelper()

	asserts.Equal(t, true, collections.Contains(sliceOfInt, 1))
	asserts.Equal(t, true, collections.Contains(sliceOfInt, 2))
	asserts.Equal(t, true, collections.Contains(sliceOfInt, 3))
	asserts.Equal(t, false, collections.Contains(sliceOfInt, 4))

	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 1, b: "a"}))
	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 2, b: "b"}))
	asserts.Equal(t, true, collections.Contains(sliceOfStruct, TestStruct{A: 3, b: "c"}))
	asserts.Equal(t, false, collections.Contains(sliceOfStruct, TestStruct{A: 4, b: "d"}))
}

func TestSlicesIndexOf(t *testing.T) {
	sliceOfInt := buildSliceOfIntsHelper()
	sliceOfStruct := buildSliceOfStructsHelper()

	asserts.Equal(t, 0, collections.IndexOf(sliceOfInt, 1))
	asserts.Equal(t, 1, collections.IndexOf(sliceOfInt, 2))
	asserts.Equal(t, 2, collections.IndexOf(sliceOfInt, 3))
	asserts.Equal(t, -1, collections.IndexOf(sliceOfInt, 4))

	asserts.Equal(t, 0, collections.IndexOf(sliceOfStruct, TestStruct{A: 1, b: "a"}))
	asserts.Equal(t, 1, collections.IndexOf(sliceOfStruct, TestStruct{A: 2, b: "b"}))
	asserts.Equal(t, 2, collections.IndexOf(sliceOfStruct, TestStruct{A: 3, b: "c"}))
	asserts.Equal(t, -1, collections.IndexOf(sliceOfStruct, TestStruct{A: 4, b: "d"}))
}

func TestSlicesRemove(t *testing.T) {
	sliceOfInt := buildSliceOfIntsHelper()
	sliceOfStruct := buildSliceOfStructsHelper()

	asserts.Slices(t, []int{2, 3}, collections.Remove(sliceOfInt, 1))
	asserts.Slices(t, []int{1, 3}, collections.Remove(sliceOfInt, 2))
	asserts.Slices(t, []int{1, 2}, collections.Remove(sliceOfInt, 3))
	asserts.Slices(t, []int{1, 2, 3}, collections.Remove(sliceOfInt, 4))

	asserts.Slices(t, []TestStruct{{A: 2, b: "b"}, {A: 3, b: "c"}}, collections.Remove(sliceOfStruct, TestStruct{A: 1, b: "a"}))
	asserts.Slices(t, []TestStruct{{A: 1, b: "a"}, {A: 3, b: "c"}}, collections.Remove(sliceOfStruct, TestStruct{A: 2, b: "b"}))
	asserts.Slices(t, []TestStruct{{A: 1, b: "a"}, {A: 2, b: "b"}}, collections.Remove(sliceOfStruct, TestStruct{A: 3, b: "c"}))
	asserts.Slices(t, []TestStruct{{A: 1, b: "a"}, {A: 2, b: "b"}, {A: 3, b: "c"}}, collections.Remove(sliceOfStruct, TestStruct{A: 4, b: "d"}))
}

func buildSliceOfIntsHelper() []int {
	return []int{1, 2, 3}
}

func buildSliceOfStructsHelper() []TestStruct {
	return []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
		{A: 3, b: "c"},
	}
}
