package main

import (
	"fmt"
)

type errorTest struct{ text string }

func main() {
	a := errorTest{"EOF"}
	b := errorTest{"EOF"}
	fmt.Println(a == b)

	c := &errorTest{"EOF"}
	d := &errorTest{"EOF"}
	fmt.Println(c == d)
}
