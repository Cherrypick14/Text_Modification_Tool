package functions

import (
	"regexp"
)

func Punc(s string) string {
	// Define the regexp for the punctuations

	// puncregexp := regexp.MustCompile(`\s+([.,!?;:]+)`)
	combinedPattern := regexp.MustCompile(`\s*([.,!?;:]+)\s*(\w*)`)

	// Replace punctuation with spaces around them
	s = combinedPattern.ReplaceAllString(s, "$1 $2")

	return s
}
