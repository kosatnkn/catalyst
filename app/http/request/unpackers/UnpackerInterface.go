package unpackers

// UnpackerInterface is the interface implemented by all unpacker data structures.
type UnpackerInterface interface {

	// RequiredFormat string representation of the required format of the relevant request body.
	RequiredFormat() string
}
