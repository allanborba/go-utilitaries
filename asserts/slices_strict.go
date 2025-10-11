package asserts

import (
	"fmt"
	"strings"
)

func SliceStrict[T comparable](t Tester, expected, actual []T) {
	slices := printableSlices(expected, actual)

	if len(expected) != len(actual) {
		t.Errorf("%s expected size of %v, received size of %v", slices, len(expected), len(actual))
	} else {
		wrongElementsIndex := getWrongElements(expected, actual)

		if wrongElementsIndex != "" {
			t.Errorf("%s diff at index %v", slices, wrongElementsIndex)
		}
	}
}

func getWrongElements[T comparable](expected []T, actual []T) string {
	wrongElementsIndex := []string{}

	for i := range expected {
		if expected[i] != actual[i] {
			stringifyedIndex := StringifyedStruct(i)
			wrongElementsIndex = append(wrongElementsIndex, stringifyedIndex)
		}
	}

	return strings.Join(wrongElementsIndex, ", ")
}

func printableSlices[T comparable](expected []T, actual []T) string {
	slices := fmt.Sprintf("\n expected %v\n got %v\n", expected, actual)
	return slices
}
