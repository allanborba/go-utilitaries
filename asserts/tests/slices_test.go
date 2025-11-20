package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestSlices_WhenHasDiffSize_ThenShowDiffSizeMsg(t *testing.T) {
	mokingT := NewFakeT()
	expected := []int{1, 2, 3}
	result := []int{1, 2}

	asserts.Slices(mokingT, expected, result)

	assertErrorMsg(t, mokingT, "expected 3 elements, got 2 elements")
}

func TestSlices_WhenAreEqualWithDiffSort_ThenShowNoError(t *testing.T) {
	mokingT := NewFakeT()
	expected := []int{1, 2, 3}

	result := []int{3, 2, 1}

	asserts.Slices(mokingT, expected, result)

	assertNoError(t, mokingT)
}

func TestSlices_WhenHasDiffElement_ThenShowDiffElementMsg(t *testing.T) {
	mokingT := NewFakeT()
	expected := []int{1, 2, 3}
	result := []int{3, 2, 4}

	asserts.Slices(mokingT, expected, result)

	assertErrorMsg(t, mokingT, "expected element 1 found 1 times, got 0 time")
}

func TestSlices_WhenSliceOfStructsHasDiffElement_ThenShowDiffElementMsg(t *testing.T) {
	mokingT := NewFakeT()
	expected := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
		{A: 3, b: "b"},
		{A: 4, b: "b"},
		{A: 5, b: "b"},
		{A: 6, b: "b"},
	}
	result := []TestStruct{
		{A: 2, b: "b"},
		{A: 4, b: "Z"},
		{A: 1, b: "a"},
		{A: 3, b: "b"},
		{A: 6, b: "b"},
		{A: 5, b: "b"},
	}

	asserts.Slices(mokingT, expected, result)

	assertErrorMsg(t, mokingT, "expected element { A: 4 b: b } found 1 times, got 0 time")
}

func TestSlices_WhenAreEqualStructsWithDiffSort_ThenShowNoError(t *testing.T) {
	mokingT := NewFakeT()
	expected := []TestStruct{
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
		{A: 1, b: "a"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
	}

	asserts.Slices(mokingT, expected, result)

	assertNoError(t, mokingT)
}

func TestSlices_WhenSliceHasMoreThaOneTimeSameElement(t *testing.T) {
	mokingT := NewFakeT()
	expected := []int{1, 2, 2, 3}
	result := []int{1, 2, 3, 3}

	asserts.Slices(mokingT, expected, result)

	assertErrorMsg(t, mokingT, "expected element 3 found 1 times, got 2 time")
}

func TestSlices_WhenSliceHasMoreThaOneTimeSameElementOfStruct(t *testing.T) {
	mokingT := NewFakeT()
	expected := []TestStruct{
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 1, b: "a"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
		{A: 4, b: "a"},
	}

	asserts.Slices(mokingT, expected, result)

	assertErrorMsg(t, mokingT, "expected element { A: 2 b: c c: { A: 3 b: 2 } } found 2 times, got 1 time")
}

func TestSlicesIgnoringFields(t *testing.T) {
	mokingT := NewFakeT()
	expected := []TestStruct{
		{A: 1, b: "1"},
		{A: 1, b: "1"},
	}
	result := []TestStruct{
		{A: 1, b: "1"},
		{A: 1, b: "2"},
	}

	asserts.SlicesIgnoringFields(mokingT, expected, result, []string{"b"})

	assertNoError(t, mokingT)
}

func TestSlicesIgnoringFields2(t *testing.T) {
	mokingT := NewFakeT()
	expected := []TestStruct{
		{A: 1, b: "1"},
		{A: 1, b: "1"},
	}
	result := []TestStruct{
		{A: 1, b: "1"},
		{A: 2, b: "2"},
	}

	asserts.SlicesIgnoringFields(mokingT, expected, result, []string{"b"})

	assertErrorMsg(t, mokingT, "expected element { A: 1 b: 1 } found 2 times, got 1 time")
}
