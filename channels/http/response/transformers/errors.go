package transformers

import (
	"fmt"

	"github.com/kosatnkn/catalyst/channels/http/errors"
)

// NOTE: Contains all common TransformerErrors

// unknownDataTypeError returns a TransformerError saying that the type of the provided data
// was not the data type expected by the transformer.
func unknownDataTypeError(expected string) error {
	return errors.NewTransformerError(fmt.Sprintf("Unknown data type found. Expect data of type %s", expected), 100, "")
}
