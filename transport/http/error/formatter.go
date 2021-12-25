package error

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kosatnkn/catalyst/v2/transport/http/response/transformers"
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

	details := strings.Split(err.Error(), "|")

	return transformers.ErrorTransformer{
		Type: details[0],
		Code: details[1],
		Msg:  details[2],
	}
}

// formatUnpackerError formats request payload unpacking errors.
//
// These occur when the format of the sent data structure does not match the expected format.
// An UnpackerError is a type of ValidationError.
func formatUnpackerError(err error) transformers.ValidationErrorTransformer {

	details := strings.Split(err.Error(), "|")

	return transformers.ValidationErrorTransformer{
		Type: "Validation Errors",
		Msg:  details[2],
	}
}

// formatValidatorErrors formats errors thrown by the validator.
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

	for i, msg := range trace {
		if i == 0 {
			t = formatForLog(msg)
			continue
		}

		t += ", " + formatForLog(msg)
	}

	return t
}

// formatForLog formats the error message for logging.
func formatForLog(msg string) string {
	return strings.Join(strings.Split(msg, "|"), " ")
}
