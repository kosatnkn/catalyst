package validator

import (
	"reflect"
)

// isSlice checks whether the given unpacker is a slice or an array.
func isSlice(data interface{}) bool {
	value := reflect.ValueOf(data)

	// if data is a pointer, get the underlying value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check if the underlying type of the value is an array
	return value.Kind() == reflect.Array || value.Kind() == reflect.Slice
}

// convertToSlice converts the given interface type that has a slice as its underlying type in to a slice of interfaces.
func convertToSlice(data interface{}) []interface{} {
	value := reflect.ValueOf(data)

	// if data is a pointer, get the underlying value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Slice {
		return nil
	}

	length := value.Len()
	result := make([]interface{}, length)

	for i := 0; i < length; i++ {
		result[i] = value.Index(i).Interface()
	}

	return result
}
