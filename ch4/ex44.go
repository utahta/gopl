package main

import (
	"fmt"
)

func rotate(s []int, i int) {
	tmp := make([]int, len(s))
	copy(tmp, s[i:])
	copy(tmp[len(s)-i:], s[:i])
	copy(s, tmp)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}
