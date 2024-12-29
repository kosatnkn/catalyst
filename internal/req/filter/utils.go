package filter

import "reflect"

// FilterByName returns the filter matching the given name.
func FilterByName(fts []Filter, name string) (Filter, bool) {
	for _, f := range fts {
		if f.Name == name {
			return f, true
		}
	}

	return Filter{}, false
}

// RemoveFilterByName removes the filter with the given name and returns the rest of the filters slice.
// This will also remove duplicate filters under the given name if there are any.
func RemoveFilterByName(fts []Filter, name string) []Filter {
	var nf []Filter

	for _, f := range fts {
		if f.Name != name {
			nf = append(nf, f)
		}
	}

	return nf
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
