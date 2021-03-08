package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kosatnkn/catalyst/v2/app/adapters"
	errHandler "github.com/kosatnkn/catalyst/v2/channels/http/error"
)

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, code int, payload []interface{}) {

	write(w, code, mapData(payload))
}

// Error formats and sends the error response.
func Error(ctx context.Context, w http.ResponseWriter, log adapters.LogAdapterInterface, err interface{}) {

	var msg interface{}
	var code int = http.StatusInternalServerError

	// check whether err is a general error or a validation error
	errG, isG := err.(error)
	errV, isV := err.(map[string]string)

	if isG {
		msg, code = errHandler.Handle(ctx, errG, log)
	}

	if isV {
		msg, code = errHandler.HandleValidatorErrors(ctx, errV, log)
	}

	write(w, code, mapErr(msg))
}

// toJSON converts the payload to JSON
func toJSON(payload interface{}) []byte {

	msg, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("JSON Marshalling Error: %v", err)
	}

	return msg
}

// write sets all required fields and write the response.
func write(w http.ResponseWriter, code int, payload interface{}) {

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// set response code
	w.WriteHeader(code)

	// set payload
	w.Write(toJSON(payload))
}
