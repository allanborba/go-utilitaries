package asserts

import "testing"

func ValidationError[T error](t *testing.T, expect string) {
	if rec := recover(); rec != nil {
		if err, ok := rec.(T); ok {
			result := err.Error()

			Equal(t, expect, result)
		} else {
			panic(rec)
		}
	} else {
		t.Errorf("expect panic, got nothing")
	}
}
