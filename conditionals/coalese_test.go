package conditionals_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
	"github.com/allanborba/utilitaries/conditionals"
)

type fakeStruct struct {
	Val int
}

func TestCoalesce(t *testing.T) {
	val1 := &fakeStruct{1}
	val2 := &fakeStruct{2}

	asserts.DeepEqual(t, &fakeStruct{1}, conditionals.Coalesce(val1))
	asserts.DeepEqual(t, &fakeStruct{2}, conditionals.Coalesce(nil, val2))
}
