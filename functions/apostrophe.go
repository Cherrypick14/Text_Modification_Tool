package functions

import (
	"strings"
)

func Apostrophe(s string) string {
	s = strings.ReplaceAll(s, "' ", " '")
	s = strings.ReplaceAll(s, " '", "'")

	return strings.TrimSpace(s)
}
