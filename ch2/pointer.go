package main

import (
	"fmt"
)

func point(p *int) {
	*p++
}

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 10
	fmt.Println(x)
	point(&x)
	fmt.Println(x)
}
