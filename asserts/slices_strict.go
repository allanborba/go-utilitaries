package asserts

import (
	"fmt"
	"reflect"
	"strings"
)

func SliceStrict[T any](t Tester, expected, result []T) {
	slices := printableSlices(expected, result)

	if len(expected) != len(result) {
		t.Errorf("%s expected size of %v, received size of %v", slices, len(expected), len(result))
	} else {
		wrongElementsIndex := getWrongElements(expected, result)

		if wrongElementsIndex != "" {
			t.Errorf("%s diff at index %v", slices, wrongElementsIndex)
		}
	}
}

func getWrongElements[T any](expected []T, actual []T) string {
	wrongElementsIndex := []string{}

	for i := range expected {
		if !reflect.DeepEqual(expected[i], actual[i]) {
			stringifyedIndex := StringifyedStruct(i)
			wrongElementsIndex = append(wrongElementsIndex, stringifyedIndex)
		}
	}

	return strings.Join(wrongElementsIndex, ", ")
}

func printableSlices[T any](expected []T, result []T) string {
	expectedString := StringifySliceOfStructs(expected)
	actualString := StringifySliceOfStructs(result)

	return fmt.Sprintf("\n expected %v\n got %v\n", expectedString, actualString)
}
