package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	for _, a := range os.Args[1:] {
		tag, e := language.Parse(a)
		if e != nil {
			fmt.Printf("%s: error: %v\n", a, e)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", a)
		} else {
			fmt.Printf("%s: tag %t\n", a, tag)
		}
	}
}
