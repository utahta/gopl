package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func diffBitsCount(a, b [32]byte) int {
	count := 0
	for i, _ := range a {
		count += int(pc[a[i]^b[i]])
	}
	return count
}

func main() {
	a := sha256.Sum256([]byte("a"))
	b := sha256.Sum256([]byte("b"))
	fmt.Println(diffBitsCount(a, b))
}
