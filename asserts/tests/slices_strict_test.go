package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestSliceStrict_WhenOneElementIsDiff_IndicateTheIndexOfDiffElement(t *testing.T) {
	mokingT := NewFakeT()
	asserts.SliceStrict(mokingT, []int{1, 2, 3}, []int{1, 2, 4})

	assertErrorMsg(t, mokingT, "\n expected [1 2 3]\n got [1 2 4]\n diff at index 2")
}

func TestSliceStrict_WhenTwoElementIsDiff_IndicateTheIndexOfDiffElement(t *testing.T) {
	mokingT := NewFakeT()
	asserts.SliceStrict(mokingT, []int{1, 2, 3, 4}, []int{2, 2, 4, 4})

	assertErrorMsg(t, mokingT, "\n expected [1 2 3 4]\n got [2 2 4 4]\n diff at index 0, 2")
}

func TestSliceStrict_WhenHasResultHasMoreElements_ShowSlicesAndSizeOfEachOne(t *testing.T) {
	mokingT := NewFakeT()
	asserts.SliceStrict(mokingT, []int{1, 2}, []int{2, 2, 4})

	assertErrorMsg(t, mokingT, "\n expected [1 2]\n got [2 2 4]\n expected size of 2, received size of 3")
}

func TestSliceStrict_WhenHasExpectedHasMoreElements_ShowSlicesAndSizeOfEachOne(t *testing.T) {
	mokingT := NewFakeT()
	asserts.SliceStrict(mokingT, []int{1, 2, 3, 4}, []int{2, 2, 4})

	assertErrorMsg(t, mokingT, "\n expected [1 2 3 4]\n got [2 2 4]\n expected size of 4, received size of 3")
}

func TestSlicesStrict_WhenElementsAreStructsAndHasOneFieldDiff_IndicateTheIndexOfDiffElement(t *testing.T) {
	mokingT := NewFakeT()
	expectedStructs := []TestStruct{
		{A: 1, b: "2"},
		{A: 5, b: "6"},
	}
	resultStructs := []TestStruct{
		{A: 1, b: "2"},
		{A: 5, b: "7"},
	}
	asserts.SliceStrict(mokingT, expectedStructs, resultStructs)

	assertErrorMsg(t, mokingT, "\n expected [{ A: 1 b: 2 } { A: 5 b: 6 }]\n got [{ A: 1 b: 2 } { A: 5 b: 7 }]\n diff at index 1")
}
