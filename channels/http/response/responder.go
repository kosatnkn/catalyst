package response

import (
	"context"
	"net/http"

	errHandler "github.com/kosatnkn/catalyst/channels/http/error"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
)

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, payload []byte, code int) {

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// set response code
	w.WriteHeader(code)

	// set payload
	w.Write(payload)
}

// Error formats and sends the error response.
func Error(ctx context.Context, w http.ResponseWriter, err interface{}, logger adapters.LogAdapterInterface) {

	var msg []byte = []byte("Unknown error type")
	var code int = http.StatusInternalServerError

	// check whether err is a general error or a validation error
	errG, isG := err.(error)
	errV, isV := err.(map[string]string)

	if isG {
		msg, code = errHandler.Handle(ctx, errG, logger)
	}

	if isV {
		msg, code = errHandler.HandleValidationErrors(ctx, errV, logger)
	}

	Send(w, msg, code)
}
