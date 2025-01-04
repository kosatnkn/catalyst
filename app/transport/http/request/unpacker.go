package request

import (
	"encoding/json"
	"regexp"

	"fmt"

	"github.com/kosatnkn/catalyst/v3/app/transport/http/request/unpacker"
)

// Unpack the request in to the given unpacker struct.
func Unpack(data []byte, un unpacker.UnpackerInterface) error {
	err := json.Unmarshal(data, un)
	if err != nil {
		return unpacker.NewUnpackerError("", formatUnpackerMessage(un.RequiredFormat()), nil)
	}

	return nil
}

// formatUnpackerMessage removes any special characters from the message string.
func formatUnpackerMessage(p string) string {
	// catch carriage returns and new lines
	reNewLine := regexp.MustCompile(`[\r\n]+`)
	// catch other special characters
	reSpecialChar := regexp.MustCompile(`[\t\"\']*`)

	m := reSpecialChar.ReplaceAllString(reNewLine.ReplaceAllString(p, " "), "")

	return fmt.Sprintf("Required format: %s", m)
}
