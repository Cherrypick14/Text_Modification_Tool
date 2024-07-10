package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// Uppercase words in a slice based on "(up)" or "(up," markers and their following numbers.

func Upper(s string) string {
	words := strings.Fields(s)
	// Iterate through each word in the slice
	for i, word := range words {
		 if word == "(up)" && i > 0 {
			prevWord := words[i-1]
			prevWord = strings.ToUpper(string(prevWord))
			words[i-1] = prevWord
			words = append(words[:i], words[i+1:]...)
		 }
		 if word == "(up," {
            trimmedText := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, err := strconv.Atoi(trimmedText)
			if err != nil {
				fmt.Println("Error converting number")
			}
			for j := 1; j <= number; j++ {
				if i-j >= 0 { // Ensure index does not go out of bounds
					prevWord := words[i-j]
					prevWord = strings.ToUpper(string(prevWord))
					words[i-j] = prevWord
				}
			}
			words = append(words[:i], words[i+2:]...)
	  		 }
		}
		return strings.Join(words, " ")
	}
	
