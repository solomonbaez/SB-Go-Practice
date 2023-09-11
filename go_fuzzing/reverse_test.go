package main

import "testing"

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
