package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	categories := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			categories["letter"]++
		} else if unicode.IsDigit(r) {
			categories["digit"]++
		}
		// ... etc ...
	}

	fmt.Printf("rune\tcount\t\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c, n := range categories {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
