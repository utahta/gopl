package main

import (
	"fmt"
)

type Hoge struct {
	foo map[int]int
}

func main() {
	h := Hoge{map[int]int{0: 1}}
	//h.foo = make(map[int]int)
	//h.foo[0] = 1
	fmt.Println(h)
}
