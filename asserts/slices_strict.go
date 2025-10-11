package asserts

import (
	"fmt"
	"strings"
)

func SliceStrict[T comparable](t Tester, expected, actual []T) {
	wrongElementsIndex := []string{}

	slices := fmt.Sprintf("\n expected %v\n got %v\n", expected, actual)

	if len(expected) != len(actual) {
		t.Errorf("%s expected size of %v, received size of %v", slices, len(expected), len(actual))
		return
	}

	for i := range expected {
		if expected[i] != actual[i] {
			stringifyedIndex := StringifyedStruct(i)
			wrongElementsIndex = append(wrongElementsIndex, stringifyedIndex)
		}
	}

	wrongElementsIndexString := strings.Join(wrongElementsIndex, ", ")

	if len(wrongElementsIndex) > 0 {
		t.Errorf("%s diff at index %v", slices, wrongElementsIndexString)
	}
}
