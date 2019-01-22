package request

import (
	"encoding/json"
	"net/http"

	"fmt"

	errTypes "github.com/kosatnkn/catalyst/app/error/types"
)

// Unpack the request in to the given unpacker struct.
func Unpack(r *http.Request, unpacker UnpackerInterface) error {

	err := json.NewDecoder(r.Body).Decode(unpacker)

	if err != nil {
		verr := errTypes.ValidationError{}

		return verr.New(fmt.Sprintf("Require following format: %s", unpacker.RequiredFormat()))
	}

	return nil
}
