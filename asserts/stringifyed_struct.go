package asserts

import (
	"fmt"
	"slices"
	"strings"

	"github.com/allanborba/go-utilitaries/collections"
)

func StringifyedStruct[T any](expected T) string {
	return StringifyedStructWithIgnoreFields(expected, []string{})
}

func StringifyedStructWithIgnoreFields[T any](expected T, fieldsToIgnore []string) string {
	if !IsStruct(expected) {
		return fmt.Sprintf("%v", expected)
	}

	mapped := StructToMap(expected)
	keys := GetFieldNames(expected)
	slices.Sort(keys)
	str := ""

	for _, k := range keys {
		v := mapped[k]
		if v == nil || IsInterfaceNil(v) || collections.Contains(fieldsToIgnore, k) {
			continue
		}

		if IsStruct(v) {
			str += fmt.Sprintf("%v: %v ", k, StringifyedStructWithIgnoreFields(v, fieldsToIgnore))
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
