package main

import (
	"fmt"
)

func hoge(a string, b []string) {
	fmt.Println(a, b)
}

func f(a []string) []string {
	return a
}

func main() {
	s := []string{"a", "b", "c"}
	a := []string{}

	a = append(a, f(s)...)
	fmt.Println(a)
}
