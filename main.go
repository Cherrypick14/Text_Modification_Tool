package main

import (
	"fmt"
	"os"

	"reloaded/functions"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: <sample.txt> <result.txt>")
		return
	}

	inputText := args[0]

	data, err := os.ReadFile(inputText)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	modifiedData := functions.Capitalize(string(data))
	modifiedData = functions.Upper(modifiedData)
	modifiedData = functions.Lower(modifiedData)
	modifiedData = functions.Vowels(modifiedData)
	modifiedData = functions.HextoDec(modifiedData)
	modifiedData = functions.BintoDec(modifiedData)
	modifiedData = functions.Punc(modifiedData)
	modifiedData = functions.Apostrophe(modifiedData)

	output := os.WriteFile(args[1], []byte(modifiedData), 0644)

	if output != nil {
		fmt.Println("err", output)
	}
}
