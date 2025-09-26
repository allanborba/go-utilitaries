package asserts

import (
	"reflect"

	"github.com/allanborba/utilitaries/collections"
)

func Slices[T comparable](t Tester, expectedSlice []T, resultSlice []T) {
	if len(expectedSlice) != len(resultSlice) {
		t.Errorf("expected %v elements, got %v elements", len(expectedSlice), len(resultSlice))
		return
	}

	resultElementsStringfyedSet := buildStringfyedSet(resultSlice)

	for _, expected := range expectedSlice {
		if IsStruct(expected) {
			stringfyedExpected := StringifyedStruct(expected)

			if !resultElementsStringfyedSet.Has(stringfyedExpected) {
				mapped := StringifyedStruct(expected)
				t.Errorf("element %v not found on results", mapped)
			}
		} else {
			if !resultElementsStringfyedSet.Has(expected) {
				t.Errorf("element %v not found on results", expected)
			}
		}
	}
}

func buildStringfyedSet[T any](resultSlice []T) *collections.Set[interface{}] {
	resultElementsStringfyedSet := collections.NewSet([]interface{}{})
	for _, result := range resultSlice {
		if IsStruct(result) {
			resultElementsStringfyedSet.Add(StringifyedStruct(result))
		} else {
			resultElementsStringfyedSet.Add(result)
		}
	}
	return resultElementsStringfyedSet
}

func _slicesSlowest[T comparable](t Tester, expectedSlice []T, resultSlice []T) {
	if len(expectedSlice) != len(resultSlice) {
		t.Errorf("expected %v elements, got %v elements", len(expectedSlice), len(resultSlice))
		return
	}

	for _, expected := range expectedSlice {
		found := false
		for _, result := range resultSlice {
			if reflect.DeepEqual(expected, result) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("element %v not found on results", StringifyedStruct(expected))
		}
	}
}
