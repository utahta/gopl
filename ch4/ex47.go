package main

import (
	"fmt"
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverse(b []byte) {
	reverseBytes(b)

	n := 0
	for i := 0; i < len(b); i++ {
		if (b[i] & 0x80) == 0 {
			continue
		}

		if (b[i] << 1 & 0x80) > 0 {
			reverseBytes(b[i-n : i+1])
			n = 0
		} else {
			n++
		}
	}
}

func main() {
	b := []byte("あいうえおabc")
	fmt.Println(string(b))
	reverse(b)
	fmt.Println(string(b))
}
