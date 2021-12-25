package unpackers

// SampleUnpacker contains the unpacking structure for sample in request payload.
//
// https://pkg.go.dev/gopkg.in/go-playground/validator.v10#section-documentation
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
