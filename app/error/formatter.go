package error

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/kosatnkn/catalyst/app/error/transformers"
	"github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/app/transport/response"
	domainError "github.com/kosatnkn/catalyst/domain/error"
)

// Format error details.
func format(err error) []byte {

	wrapper := response.Error{}
	var payload interface{}

	switch err.(type) {
	case *types.ServerError, *types.MiddlewareError, *types.AdapterError, *types.RepositoryError, *types.ServiceError, *domainError.DomainError:
		payload = formatCustomError(err)
		break
	case *types.ValidationError:
		payload = formatValidationStructureError(err)
		break
	default:
		payload = formatUnknownError(err)
		break
	}

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// Format custom errors.
func formatCustomError(err error) transformers.ErrorTransformer {

	errorDetails := strings.Split(err.Error(), "|")
	errCode, _ := strconv.Atoi(errorDetails[1])

	payload := transformers.ErrorTransformer{}
	payload.Msg = errorDetails[0]
	payload.Code = errCode
	payload.Type = errorDetails[2]
	payload.Trace = errorDetails[3]

	return payload
}

// Format validation structure errors.
// These occur when the format of the sent data structure does not match the expected format.
func formatValidationStructureError(err error) transformers.ValidationErrorTransformer {

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = err.Error()

	return payload
}

// Format validation errors.
// These are errors thrown when field wise validations happen against a data structure.
func formatValidationErrors(p map[string]string) []byte {

	wrapper := response.Error{}

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = formatValidationPayload(p)

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// Do a final round of formatting to validation errors.
func formatValidationPayload(p map[string]string) map[string]string {

	ep := make(map[string]string)

	for k, v := range p {
		ek := strcase.ToSnake(k)
		ep[ek] = v
	}

	return ep
}

// Format errors of unhandled types.
func formatUnknownError(err error) transformers.ErrorTransformer {

	payload := transformers.ErrorTransformer{}
	payload.Type = "Unknown Error"
	payload.Msg = err.Error()

	return payload
}
