package main

import (
	"fmt"
)

func removeAdjacentDup(s []string) []string {
	var tmp string
	i := 0
	for i < len(s) {
		if tmp == s[i] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			tmp = s[i]
			i++
		}
	}
	return s
}

func main() {
	a := []string{"a", "a", "a", "b", "b", "c", "c", "a", "b", "c"}
	fmt.Println(a)
	fmt.Println(removeAdjacentDup(a))
}
