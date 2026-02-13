package asserts

const BOOL_ERROR_MSG = "expected %v, got %v"

func True(t Tester, result bool) {
	if !result {
		t.Errorf(BOOL_ERROR_MSG, true, result)
	}
}

func False(t Tester, result bool) {
	if result {
		t.Errorf(BOOL_ERROR_MSG, false, result)
	}
}
