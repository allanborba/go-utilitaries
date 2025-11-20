package asserts

func Slices[T any](t Tester, expectedSlice []T, resultSlice []T) {
	SlicesIgnoringFields(t, expectedSlice, resultSlice, []string{})
}

func SlicesIgnoringFields[T any](t Tester, expectedSlice []T, resultSlice []T, ignoreFields []string) {
	if len(expectedSlice) != len(resultSlice) {
		t.Errorf("expected %v elements, got %v elements", len(expectedSlice), len(resultSlice))
		return
	}

	resultElementsStringfyedSet := buildStringfyedSet(resultSlice, ignoreFields)
	expectElementsStringfyedSet := buildStringfyedSet(expectedSlice, ignoreFields)

	for _, expected := range expectedSlice {
		expectedCount, resultCount := getElementsCount(expected, expectElementsStringfyedSet, resultElementsStringfyedSet, ignoreFields)

		if resultCount != expectedCount {
			t.Errorf("expected element %v found %d times, got %d time", StringifyedStruct(expected), expectedCount, resultCount)
		}
	}
}

func getElementsCount[T any](expected T, expectElementsStringfyedSet, resultElementsStringfyedSet map[interface{}]int, fieldsToIgnore []string) (int, int) {
	var expectedCount int
	var resultCount int

	if IsStruct(expected) {
		expectedCount = expectElementsStringfyedSet[StringifyedStructWithIgnoreFields(expected, fieldsToIgnore)]
		resultCount = resultElementsStringfyedSet[StringifyedStructWithIgnoreFields(expected, fieldsToIgnore)]
	} else {
		expectedCount = expectElementsStringfyedSet[expected]
		resultCount = resultElementsStringfyedSet[expected]
	}

	return expectedCount, resultCount
}

func buildStringfyedSet[T any](resultSlice []T, ignoreFields []string) map[interface{}]int {
	resultElementsStringfyedSet := make(map[interface{}]int)

	for _, result := range resultSlice {
		if IsStruct(result) {
			resultElementsStringfyedSet[StringifyedStructWithIgnoreFields(result, ignoreFields)]++
		} else {
			resultElementsStringfyedSet[result]++
		}
	}

	return resultElementsStringfyedSet
}
