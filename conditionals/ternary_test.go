package conditionals_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
	"github.com/allanborba/go-utilitaries/conditionals"
)

func TestTernary(t *testing.T) {
	asserts.Equal(t, 1, conditionals.Ternary(true, 1, 2))
	asserts.Equal(t, 2, conditionals.Ternary(false, 1, 2))

	asserts.Equal(t, "a", conditionals.Ternary(true, "a", "b"))
	asserts.Equal(t, "b", conditionals.Ternary(false, "a", "b"))
}
