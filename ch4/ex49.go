package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("word.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	counts := map[string]int{}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%q\t%d\n", k, v)
	}
}
