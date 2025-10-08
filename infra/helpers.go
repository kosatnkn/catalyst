package infra

import "strings"

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
