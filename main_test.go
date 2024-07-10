package main

import (
	"testing"

	"reloaded/functions"
)

func TestCapitalize(t *testing.T) {
	cap := []struct {
		input    string
		expected string
	}{
		{"it (cap)", "It"},
		{"it was the age of foolishness (cap, 6)", "It Was The Age Of Foolishness"},
		{"the sun (cap, 8) rises in the east", "The Sun rises in the east"},
		{"the sun (cap) rises in the east", "the Sun rises in the east"},
		{"the sun (capijiooimoimoi) rises in the east", "the sun (capijiooimoimoi) rises in the east"},
		{"the sun (oioimiomoimcap) rises in the east", "the sun (oioimiomoimcap) rises in the east"},
		{"the sun (oioimiomoimCAP) rises in the east", "the sun (oioimiomoimCAP) rises in the east"},
		{"the sun (oioimiomoimCaP) rises in the east", "the sun (oioimiomoimCaP) rises in the east"},
	}

	for _, char := range cap {
		output := functions.Capitalize(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s\n", char.expected, output)
		}
	}
}

func TestUpper(t *testing.T) {
	up := []struct {
		input    string
		expected string
	}{
		{"it (up)", "IT"},
		{"it was the age of foolishness (up, 3)", "it was the AGE OF FOOLISHNESS"},
		{"The Sun (up, 6) rises in the east", "THE SUN rises in the east"},
		{"the Sun (up) rises in the east", "the SUN rises in the east"},
		{"the sun (UPijiooimoimoi) rises in the east", "the sun (UPijiooimoimoi) rises in the east"},
		{"the sun (oioimiomoimup) rises in the east", "the sun (oioimiomoimup) rises in the east"},
		{"the sun (oioimiomoimuP) rises in the east", "the sun (oioimiomoimuP) rises in the east"},
		{"the sun (oioimiomoimUp) rises in the east", "the sun (oioimiomoimUp) rises in the east"},
	}

	for _, char := range up {
		output := functions.Upper(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}

func TestLower(t *testing.T) {
	low := []struct {
		input    string
		expected string
	}{
		{"IT (low)", "it"},
		{"It Was The Age Of Foolishness (low, 5)", "It was the age of foolishness"},
		{"The Sun (low, 2) rises in the east", "the sun rises in the east"},
		{"the Sun (low) rises in the east", "the sun rises in the east"},
		{"the sun (LOWijiooimoimoi) rises in the east", "the sun (LOWijiooimoimoi) rises in the east"},
		{"the sun (oioimiomoimlow) rises in the east", "the sun (oioimiomoimlow) rises in the east"},
		{"the sun (oioimiomoimLoW) rises in the east", "the sun (oioimiomoimLoW) rises in the east"},
		{"the sun (oioimiomoimloW) rises in the east", "the sun (oioimiomoimloW) rises in the east"},
	}

	for _, char := range low {
		output := functions.Lower(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}

func TestVowels(t *testing.T) {
	vowel := []struct {
		input    string
		expected string
	}{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"A apple", "An apple"},
		{"A hour", "An hour"},
	}

	for _, char := range vowel {
		output := functions.Vowels(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}

func TestApostrophe(t *testing.T) {
	apos := []struct {
		input    string
		expected string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}

	for _, char := range apos {
		output := functions.Apostrophe(char.input)
		if output != char.expected {
			t.Errorf("Expected:%s Output: %s \n", char.expected, output)
		}
	}
}

func TestPunc(t *testing.T) {
	punc := []struct {
		input    string
		expected string
	}{
		{"I was thinking ... You were right", "I was thinking... You were right"},
		
	}

	for _, char := range punc {
		output := functions.Punc(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}

func TestHextoDec(t *testing.T) {
	hex := []struct {
		input    string
		expected string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"Yesterday, a batch of 1E (hex) files were added to the system.", "Yesterday, a batch of 30 files were added to the system."},
	}

	for _, char := range hex {
		output := functions.HextoDec(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}

func TestBintoDec(t *testing.T) {
	bin := []struct {
		input    string
		expected string
	}{
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"As of today, it has been 10 (bin) years since the project began.", "As of today, it has been 2 years since the project began."},
	}

	for _, char := range bin {
		output := functions.BintoDec(char.input)
		if output != char.expected {
			t.Errorf("Expected: %s Output: %s \n", char.expected, output)
		}
	}
}
