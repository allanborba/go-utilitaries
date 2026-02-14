package asserts

import (
	"fmt"
	"reflect"
	"strings"
)

func Slices[T any](t Tester, expected, result []T) {
	missing, extra := diffSlices(expected, result)

	msgs := []string{}

	if len(missing) > 0 {
		msgs = append(msgs, fmt.Sprintf("\nmissing elements:\n%s", formatElements(missing)))
	}

	if len(extra) > 0 {
		msgs = append(msgs, fmt.Sprintf("\nextra elements:\n%s", formatElements(extra)))
	}

	if len(msgs) > 0 {
		t.Errorf(strings.Join(msgs, "\n"))
	}
}

func diffSlices[T any](expected, result []T) (missing []T, extra []T) {
	resultUsed := make([]bool, len(result))

	for _, exp := range expected {
		found := false
		for j, res := range result {
			if resultUsed[j] {
				continue
			}
			if objectsEqual(exp, res) {
				resultUsed[j] = true
				found = true
				break
			}
		}
		if !found {
			missing = append(missing, exp)
		}
	}

	for j, res := range result {
		if !resultUsed[j] {
			extra = append(extra, res)
		}
	}

	return missing, extra
}

func objectsEqual[T any](a, b T) bool {
	if IsStruct(a) {
		return structsEqual(a, b)
	}
	return reflect.DeepEqual(a, b)
}

func structsEqual[T any](a, b T) bool {
	mapA := StructToMap(a)
	mapB := StructToMap(b)
	fieldsA := GetFieldNames(a)

	for _, field := range fieldsA {
		valA := mapA[field]
		valB := mapB[field]

		if valA == nil && valB == nil {
			continue
		}
		if IsInterfaceNil(valA) && IsInterfaceNil(valB) {
			continue
		}
		if IsInterfaceNil(valA) != IsInterfaceNil(valB) {
			return false
		}

		if isSliceValue(valA) {
			if !sliceFieldsEqual(valA, valB) {
				return false
			}
		} else if IsStruct(valA) {
			if !structsEqual(valA, valB) {
				return false
			}
		} else {
			if !reflect.DeepEqual(valA, valB) {
				return false
			}
		}
	}

	return true
}

func isSliceValue(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func sliceFieldsEqual(a, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Len() != vb.Len() {
		return false
	}

	used := make([]bool, vb.Len())

	for i := 0; i < va.Len(); i++ {
		elemA := va.Index(i).Interface()
		found := false
		for j := 0; j < vb.Len(); j++ {
			if used[j] {
				continue
			}
			elemB := vb.Index(j).Interface()
			if objectsEqual(elemA, elemB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func formatElements[T any](elements []T) string {
	strs := make([]string, len(elements))
	for i, e := range elements {
		if IsStruct(e) {
			strs[i] = fmt.Sprintf("  - %s", StringifyedStruct(e))
		} else {
			strs[i] = fmt.Sprintf("  - %v", e)
		}
	}
	return strings.Join(strs, "\n")
}
