package unpackers

// SampleFiltersUnpacker contains the unpacking structure for the school `filters` query parameter.
//
// https://pkg.go.dev/gopkg.in/go-playground/validator.v10#section-documentation
type SampleFiltersUnpacker struct {
	NameContain string `json:"name_contain" validate:"omitempty"`
}

// NewSampleFilterUnpacker creates a new instance of the unpacker.
func NewSampleFiltersUnpacker() *SampleFiltersUnpacker {
	return &SampleFiltersUnpacker{}
}

// RequiredFormat returns the applicable JSON format for the school data structure.
func (u *SampleFiltersUnpacker) RequiredFormat() string {
	return `{
		"name_contain": "<string, optional>"
	}`
}
