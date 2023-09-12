package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"!12345", "54321!"},
		{" ", " "},
	}

	for _, testCase := range testCases {
		reversedString := StringReverse(testCase.in)
		if reversedString != testCase.want {
			t.Errorf("Reverse: %v, want %v", reversedString, testCase.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, testCase := range testCases {
		f.Add(testCase) // seed the fuzz
	}

	f.Fuzz(func(t *testing.T, original string) {
		singleReverse := StringReverse(original)
		doubleReverse := StringReverse(singleReverse)
		t.Logf("Number of runes: original=%d, reverse=%d, double-reverse=%d",
			utf8.RuneCountInString(original),
			utf8.RuneCountInString(singleReverse),
			utf8.RuneCountInString(doubleReverse),
		)

		if original != doubleReverse {
			t.Errorf("Before: %v, after: %v", original, doubleReverse)
		}

		if utf8.ValidString(original) && !utf8.ValidString(singleReverse) {
			t.Errorf("StringReverse produced invalid UTF-8 string %v", singleReverse)
		}
	})
}
