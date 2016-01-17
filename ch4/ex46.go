package main

import (
	"fmt"
	"unicode"
)

func squash(bytes []byte) {
	for i, b := range bytes {
		if unicode.IsSpace(rune(b)) {
			bytes[i] = ' '
		}
	}
}

func main() {
	a := []byte("a\nb\rc\td　e ")
	fmt.Println(string(a))
	squash(a)
	fmt.Println(string(a))
}
