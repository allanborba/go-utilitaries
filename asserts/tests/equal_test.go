package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestEqual_WhenValueAreEqual_ThenNotShowError(t *testing.T) {
	assertEqualNoError(t, 0, 0)
	assertEqualNoError(t, "a", "a")
	assertEqual(t, 0, 1, "expected 0, got 1")
	assertEqual(t, "0", "1", "expected 0, got 1")
	assertEqualNoError(t, true, true)
	assertEqual(t, true, false, "expected true, got false")
}

func assertEqualNoError[T any](t *testing.T, expected, actual T) {
	mockT := NewFakeT()
	asserts.Equal(mockT, expected, actual)
	assertNoError(t, mockT)
}

func assertEqual[T any](t *testing.T, expected, actual T, msg string) {
	mockT := NewFakeT()
	asserts.Equal(mockT, expected, actual)
	assertErrorMsg(t, mockT, msg)
}
