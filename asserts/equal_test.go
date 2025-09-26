package asserts_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
)

func TestEqual_WhenValueAreEqual_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 0, 0)

	assertNoError(t, mokingT)
}

func TestEqual_WhenValueAreNotEqual_ThenShowNotEqualErrorMsg(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 1, 0)

	assertErrorMsg(t, mokingT, "expected 1, got 0")
}
