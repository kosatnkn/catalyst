package request

// UnpackerInterface is the interface implemented by all unpacker data structures.
type UnpackerInterface interface {

	// String representation of the required format of the relevant request body.
	RequiredFormat() string
}
