package unpackers

// SampleUnpacker contains the unpacking structure for the address sent in request payload.
type SampleUnpacker struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// NewSampleUnpacker creates a new instance of the unpacker.
func NewSampleUnpacker() *SampleUnpacker {

	return &SampleUnpacker{}
}

// RequiredFormat returns the applicable JSON format for the address data structure.
func (u *SampleUnpacker) RequiredFormat() string {

	return `{
		"name": "<string>",
		"password": "<string>"
	}`
}
