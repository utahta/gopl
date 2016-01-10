package main

import (
	"fmt"
)

func main() {
	a, b, c, d := 1, 2, 3, 4
	fmt.Println(a, b, c, d)
	a, b, c, d = d, c, b, a
	fmt.Println(a, b, c, d)
}
