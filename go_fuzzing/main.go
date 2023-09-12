package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "stinky"
	singleReverse, singleReverseError := StringReverse(s)
	doubleReverse, doubleReverseError := StringReverse(singleReverse)
	fmt.Printf("original: %v\n", s)
	fmt.Printf("reversed: %v, error: %v\n", singleReverse, singleReverseError)
	fmt.Printf("reversed: %v, error: %v\n", doubleReverse, doubleReverseError)
}

func StringReverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r), nil
}
