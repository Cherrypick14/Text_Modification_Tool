package functions

import (
	"strings"
)

// The Vowels function takes a string, splits it and ranges through it checking if the first char of

// the next word is a vowel, if it is then it changes the string "a" to "an".

// This function also handles cases where the first character of the next word is a consonant and

// if it is, it changes the string "an" back to "a". Also if the string is "an" and the first char

// of the next word is a vowel, it retains the original string.

func Vowels(s string) string {
	words := strings.Fields(s)
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(words); i++ {
		for _, char := range vowels {

			if words[i] == "an" && string(words[i+1][0]) == char {
				words[i] = "an"
				break
			}
			if words[i] == "An" && string(words[i+1][0]) == char {
				break
			}
			if words[i] == "A" && string(words[i+1][0]) == char {
				words[i] = "An"
				break
			}
			if words[i] == "An" && string(words[i+1][0]) != char {
				words[i] = "A"
			}
			if words[i] == "a" && string(words[i+1][0]) == char {
				words[i] = "an"
			}

			if words[i] == "an" && string(words[i+1][0]) != char {
				words[i] = "a"
			}
		}
	}
	return strings.Join(words, " ")
}
