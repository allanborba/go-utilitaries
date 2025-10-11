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

func TestSliceStrict_WhenHasDiffSize_ShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.SliceStrict(mokingT, []int{1, 2}, []int{2, 2, 4})

	assertErrorMsg(t, mokingT, "\n expected [1 2]\n got [2 2 4]\n expected size of 2, received size of 3")
}
