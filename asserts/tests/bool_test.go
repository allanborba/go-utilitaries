package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestTrue_WhenValueIsTrue_ThenNotShowError(t *testing.T) {
	mockT := NewFakeT()
	asserts.True(mockT, true)
	assertNoError(t, mockT)
}

func TestTrue_WhenValueIsFalse_ThenShowError(t *testing.T) {
	mockT := NewFakeT()
	asserts.True(mockT, false)
	assertErrorMsg(t, mockT, "expected true, got false")
}

func TestFalse_WhenValueIsFalse_ThenNotShowError(t *testing.T) {
	mockT := NewFakeT()
	asserts.False(mockT, false)
	assertNoError(t, mockT)
}

func TestFalse_WhenValueIsTrue_ThenShowError(t *testing.T) {
	mockT := NewFakeT()
	asserts.False(mockT, true)
	assertErrorMsg(t, mockT, "expected false, got true")
}
