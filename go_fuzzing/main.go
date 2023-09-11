package main

import "fmt"

func main() {
	s := "stinky"
	singleReverse := StringReverse(s)
	doubleReverse := StringReverse(singleReverse)
	fmt.Printf("%v is %v backwards!\n", singleReverse, s)
	fmt.Printf("%v is %v backwards!\n", doubleReverse, singleReverse)
}

func StringReverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
