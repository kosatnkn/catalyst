package infra

import "strings"

// FormatMsg to make it JSON friendly.
func FormatMsg(msg string) string {
	sChrs := map[string]string{
		"\n": ",",
		"\r": ",",
		"\t": " ",
	}

	msg = strings.TrimSpace(msg)
	for k, v := range sChrs {
		msg = strings.ReplaceAll(msg, k, v)
	}

	return msg
}
