package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func BintoDec(s string) string {
	// Split the input string into words
	words := strings.Fields(s)
	// Iterate over the words
	for i, word := range words {
		// Check if the current word is "(bin)"
		if word == "(bin)" {
			// If it is, get the previous word
			prevWord := words[i-1]
			// Convert the previous word from binary to decimal
			number, err := strconv.ParseInt(string(prevWord), 2, 64)
			if err != nil {
				// If conversion fails, return an empty string
				return ""
			}
			// Replace the previous word with its decimal equivalent
			words[i-1] = fmt.Sprint((number))
			// Remove the "(bin)" word from the slice
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
