package asserts

import (
	"reflect"
	"unsafe"
)

func GetFieldNames[T any](obj T) []string {
	fields := []string{}
	value := reflect.ValueOf(obj)
	typeOf := value.Type()

	for i := range typeOf.NumField() {
		filedName := typeOf.Field(i).Name
		fields = append(fields, filedName)
	}

	return fields
}

func StructToMap(in interface{}) map[string]interface{} {
	structMapped := make(map[string]interface{})

	value := reflect.ValueOf(in)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	copy := reflect.New(value.Type()).Elem()
	copy.Set(value)

	elType := copy.Type()

	for i := 0; i < copy.NumField(); i++ {
		mapValue(copy, i, elType, structMapped)
	}

	return structMapped
}

func mapValue(copy reflect.Value, i int, elType reflect.Type, structMapped map[string]interface{}) {
	field := copy.Field(i)
	fieldName := elType.Field(i).Name

	accessibleField := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()

	structMapped[fieldName] = accessibleField.Interface()
}
