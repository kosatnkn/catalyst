package error

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"

	baseErrs "github.com/kosatnkn/catalyst/app/errors"
	httpErrs "github.com/kosatnkn/catalyst/app/http/errors"
	"github.com/kosatnkn/catalyst/app/http/response/mappers"
	"github.com/kosatnkn/catalyst/app/http/response/transformers"
	domainErrs "github.com/kosatnkn/catalyst/domain/error"
	externalErrs "github.com/kosatnkn/catalyst/externals/errors"
)

// format formats the error by error type.
func format(err error) []byte {

	var payload interface{}

	switch err.(type) {
	case *baseErrs.ServerError,
		*httpErrs.MiddlewareError,
		*externalErrs.AdapterError,
		*externalErrs.RepositoryError,
		*externalErrs.ServiceError,
		*domainErrs.DomainError:
		payload = formatGenericError(err)
		break
	case *httpErrs.ValidationError:
		payload = formatUnpackerError(err)
		break
	default:
		payload = formatUnknownError(err)
		break
	}

	wrapper := mappers.Error{
		Payload: payload,
	}

	message, _ := json.Marshal(wrapper)

	return message
}

// formatGenericError formats all generic errors.
func formatGenericError(err error) transformers.ErrorTransformer {

	errorDetails := strings.Split(err.Error(), "|")
	errCode, _ := strconv.Atoi(errorDetails[1])

	return transformers.ErrorTransformer{
		Type:  errorDetails[0],
		Code:  errCode,
		Msg:   errorDetails[2],
		Trace: errorDetails[3],
	}
}

// formatUnpackerError formats request payload unpacking errors.
//
// These occur when the format of the sent data structure does not match the expected format.
// An UnpackerError is a type of ValidationError.
func formatUnpackerError(err error) transformers.ValidationErrorTransformer {

	return transformers.ValidationErrorTransformer{
		Type:  "Validation Errors",
		Trace: err.Error(),
	}
}

// formatValidationErrors formats validation errors.
//
// These are errors thrown when field wise validations of the data structure fails.
func formatValidationErrors(p map[string]string) []byte {

	payload := transformers.ValidationErrorTransformer{
		Type:  "Validation Errors",
		Trace: formatValidationPayload(p),
	}

	wrapper := mappers.Error{
		Payload: payload,
	}

	message, _ := json.Marshal(wrapper)

	return message
}

// formatUnknownError formats errors of unknown error types.
func formatUnknownError(err error) transformers.ErrorTransformer {

	return transformers.ErrorTransformer{
		Type: "Unknown Error",
		Msg:  err.Error(),
	}
}

// formatValidationPayload does a final round of formatting to validation errors.
func formatValidationPayload(p map[string]string) map[string]string {

	ep := make(map[string]string)

	for k, v := range p {

		ek := formatKey(k)
		ep[ek] = v
	}

	return ep
}

// formatKey formats the key as a snake case string consisting only of lowecase characters.
func formatKey(k string) string {

	kParts := strings.Split(k, ".")

	// remove unpacker name
	kParts = kParts[1:]

	for i, part := range kParts {
		kParts[i] = strcase.ToSnake(part)
	}

	return strings.Join(kParts, ".")
}
