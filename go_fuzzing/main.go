package main

import "fmt"

func main() {
	s := "stinky"
	fmt.Printf("%v is %v backwards!\n", StringReverse(s), s)
}

func StringReverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
