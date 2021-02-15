package error

import (
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kosatnkn/catalyst/channels/http/response/transformers"
)

// formatUnknownError formats errors of unknown error types.
func formatUnknownError(err error) transformers.ErrorTransformer {

	return transformers.ErrorTransformer{
		Type: "Unknown Error",
		Msg:  err.Error(),
	}
}

// formatGenericError formats all generic errors.
func formatGenericError(err error) transformers.ErrorTransformer {

	errorDetails := strings.Split(err.Error(), "|")
	errCode, _ := strconv.Atoi(errorDetails[1])

	return transformers.ErrorTransformer{
		Type: errorDetails[0],
		Code: errCode,
		Msg:  errorDetails[2],
	}
}

// formatValidationError formats request payload unpacking errors.
//
// These occur when the format of the sent data structure does not match the expected format.
// An UnpackerError is a type of ValidationError.
func formatValidationError(err error) transformers.ValidationErrorTransformer {

	return transformers.ValidationErrorTransformer{
		Type: "Validation Errors",
		Msg:  err.Error(),
	}
}

// formatValidatorErrors formats validation errors.
//
// These are errors thrown when field wise validations of the data structure fails.
func formatValidatorErrors(p map[string]string) transformers.ValidationErrorTransformer {

	return transformers.ValidationErrorTransformer{
		Type: "Validation Errors",
		Msg:  formatValidationPayload(p),
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

// formatLogTrace formats tracing information for logging.
func formatLogTrace(trace []string) (t string) {

	for _, msg := range trace {
		t += " >> " + formatForLog(msg)
	}

	return t
}

// formatForLog formats the error message for logging.
func formatForLog(msg string) string {
	return strings.Join(strings.Split(msg, "|"), " ")
}
