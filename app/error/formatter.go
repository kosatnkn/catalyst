package error

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/kosatnkn/catalyst/app/error/transformers"
	"github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/app/transport/response/mappers"
	domainError "github.com/kosatnkn/catalyst/domain/error"
)

// format formats the error by error type.
func format(err error) []byte {

	wrapper := mappers.Error{}
	var payload interface{}

	switch err.(type) {
	case *types.ServerError,
		*types.MiddlewareError,
		*types.AdapterError,
		*types.RepositoryError,
		*types.ServiceError,
		*domainError.DomainError:
		payload = formatCustomError(err)
		break
	case *types.ValidationError:
		payload = formatUnpackerError(err)
		break
	default:
		payload = formatUnknownError(err)
		break
	}

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// formatCustomError formats all generic errors.
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

// formatUnpackerError formats request payload unpacking errors.
// These occur when the format of the sent data structure does not match the expected format.
// An UnpackerError is a type of ValidationError.
func formatUnpackerError(err error) transformers.ValidationErrorTransformer {

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = err.Error()

	return payload
}

// formatValidationErrors formats validation errors.
// These are errors thrown when field wise validations of the data structure fails.
func formatValidationErrors(p map[string]string) []byte {

	wrapper := mappers.Error{}

	payload := transformers.ValidationErrorTransformer{}
	payload.Type = "Validation Errors"
	payload.Trace = formatValidationPayload(p)

	wrapper.Payload = payload
	message, _ := json.Marshal(wrapper)

	return message
}

// formatUnknownError formats errors of unknown error types.
func formatUnknownError(err error) transformers.ErrorTransformer {

	payload := transformers.ErrorTransformer{}
	payload.Type = "Unknown Error"
	payload.Msg = err.Error()

	return payload
}

// formatValidationPayload does a final round of formatting to validation errors.
func formatValidationPayload(p map[string]string) map[string]string {

	ep := make(map[string]string)

	for k, v := range p {
		ek := strcase.ToSnake(k)
		ep[ek] = v
	}

	return ep
}
