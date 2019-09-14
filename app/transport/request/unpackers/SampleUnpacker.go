package unpackers

// SampleUnpacker contains the unpacking structure for the address sent in request payload.
type SampleUnpacker struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

// NewSampleUnpacker creates a new instance of the unpacker.
func NewSampleUnpacker() *SampleUnpacker {

	return &SampleUnpacker{}
}

// RequiredFormat returns the applicable JSON format for the address data structure.
func (u *SampleUnpacker) RequiredFormat() string {

	return `{
		"street": <string>,
		"city": <string>,
		"planet": <string>,
		"phone": <string>
	}`
}
