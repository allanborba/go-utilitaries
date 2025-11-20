package asserts

func Slices[T any](t Tester, expectedSlice []T, resultSlice []T) {
	if len(expectedSlice) != len(resultSlice) {
		t.Errorf("expected %v elements, got %v elements", len(expectedSlice), len(resultSlice))
		return
	}

	resultElementsStringfyedSet := buildStringfyedSet(resultSlice)
	expectElementsStringfyedSet := buildStringfyedSet(expectedSlice)

	for _, expected := range expectedSlice {
		expectedCount, resultCount := getElementsCount(expected, expectElementsStringfyedSet, resultElementsStringfyedSet)

		if resultCount != expectedCount {
			t.Errorf("expected element %v found %d times, got %d time", StringifyedStruct(expected), expectedCount, resultCount)
		}
	}
}

func getElementsCount[T any](expected T, expectElementsStringfyedSet, resultElementsStringfyedSet map[interface{}]int) (int, int) {
	var expectedCount int
	var resultCount int

	if IsStruct(expected) {
		expectedCount = expectElementsStringfyedSet[StringifyedStruct(expected)]
		resultCount = resultElementsStringfyedSet[StringifyedStruct(expected)]
	} else {
		expectedCount = expectElementsStringfyedSet[expected]
		resultCount = resultElementsStringfyedSet[expected]
	}

	return expectedCount, resultCount
}

func buildStringfyedSet[T any](resultSlice []T) map[interface{}]int {
	resultElementsStringfyedSet := make(map[interface{}]int)

	for _, result := range resultSlice {
		if IsStruct(result) {
			resultElementsStringfyedSet[StringifyedStruct(result)]++
		} else {
			resultElementsStringfyedSet[result]++
		}
	}

	return resultElementsStringfyedSet
}
