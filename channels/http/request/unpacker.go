package request

import (
	"encoding/json"
	"regexp"

	"fmt"

	"github.com/kosatnkn/catalyst/v2/channels/http/request/unpackers"
	"github.com/kosatnkn/catalyst/v2/channels/http/request/unpackers/errors"
)

// Unpack the request in to the given unpacker struct.
func Unpack(data []byte, unpacker unpackers.UnpackerInterface) error {

	err := json.Unmarshal(data, unpacker)
	if err != nil {
		return errors.NewUnpackerError("", formatUnpackerMessage(unpacker.RequiredFormat()), nil)
	}

	return nil
}

// formatUnpackerMessage removes any special chatacters from the message string.
func formatUnpackerMessage(p string) string {

	// catch carrage returns and new lines
	reNewLine := regexp.MustCompile(`[\r\n]+`)

	// catch other special characters
	reSpecialChar := regexp.MustCompile(`[\t\"\']*`)

	m := reSpecialChar.ReplaceAllString(reNewLine.ReplaceAllString(p, " "), "")

	return fmt.Sprintf("Required format: %s", m)
}
