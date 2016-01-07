package main

import (
	"fmt"
)

func localvar() *int {
	a := 10
	fmt.Println(&a)
	return &a
}

func main() {
	var b *int = localvar()
	fmt.Println(b)
}
