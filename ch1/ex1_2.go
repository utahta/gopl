package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s%d:%s", sep, i, arg)
		sep = " "
	}
	fmt.Println(s)
}
