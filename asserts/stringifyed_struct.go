package asserts

import (
	"fmt"
	"slices"
	"strings"
)

func StringifyedStruct[T any](expected T) string {
	if !IsStruct(expected) {
		return fmt.Sprintf("%v", expected)
	}

	mapped := StructToMap(expected)
	keys := GetFieldNames(expected)
	slices.Sort(keys)
	str := ""

	for _, k := range keys {
		v := mapped[k]
		if v == nil || IsInterfaceNil(v) {
			continue
		}

		if IsStruct(v) {
			str += fmt.Sprintf("%v: %v ", k, StringifyedStruct(v))
		} else {
			str += fmt.Sprintf("%v: %v ", k, v)
		}
	}

	return fmt.Sprint("{ ", str, "}")
}

func StringifySliceOfStructs[T any](slice []T) string {
	sliceofStrings := []string{}
	for _, item := range slice {
		sliceofStrings = append(sliceofStrings, StringifyedStruct(item))
	}

	return fmt.Sprintf("[%s]", strings.Join(sliceofStrings, " "))
}
