package filter

import (
	"reflect"
)

// FilterControllerFacilitator is the facilitator that will add filter handling capabilities to the controller.
type FilterControllerFacilitator struct{}

// NewFilterControllerFacilitator creates a new instance of the facilitator.
func NewFilterControllerFacilitator() *FilterControllerFacilitator {
	return &FilterControllerFacilitator{}
}

// GetFilters return the data passed in to it as a slice of filters.
//
// The underlying type of 'data' should be a slice.
func (ctl *FilterControllerFacilitator) GetFilters(data interface{}) (filters []Filter, err error) {
	fts := convertToSlice(data)
	for _, ft := range fts {
		f, err := ctl.mapFilter(ft)
		if err != nil {
			return filters, err
		}

		filters = append(filters, f)
	}

	return filters, nil
}

// mapFilter maps the ftr unpacker to filter entity.
func (ctl *FilterControllerFacilitator) mapFilter(ftr interface{}) (filter Filter, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	value := reflect.ValueOf(ftr)
	f := Filter{
		Name:  value.FieldByName("Name").String(),
		Value: value.FieldByName("Value").Interface(),
	}

	return f, nil
}
