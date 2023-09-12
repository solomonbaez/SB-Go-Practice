package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, testCase := range testCases {
		f.Add(testCase) // seed the fuzz
	}

	f.Fuzz(func(t *testing.T, original string) {
		singleReverse, e_1 := StringReverse(original)
		if e_1 != nil {
			return
		}
		doubleReverse, e_2 := StringReverse(singleReverse)
		if e_2 != nil {
			return
		}

		if original != doubleReverse {
			t.Errorf("Before: %v, after: %v", original, doubleReverse)
		}

		if utf8.ValidString(original) && !utf8.ValidString(singleReverse) {
			t.Errorf("StringReverse produced invalid UTF-8 string %v", singleReverse)
		}
	})
}
