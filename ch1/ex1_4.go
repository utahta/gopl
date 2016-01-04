package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counts struct {
	count     int
	filenames map[string]bool
}

func countLines(f *os.File, counts map[string]*Counts) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		_, ok := counts[text]
		if ok {
			counts[text].count++
			counts[text].filenames[f.Name()] = true
		} else {
			counts[text] = &Counts{1, map[string]bool{f.Name(): true}}
		}
	}
}

func main() {
	counts := make(map[string]*Counts)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, c := range counts {
		if c.count > 1 {
			fmt.Printf("%d:%v\t%s\n", c.count, c.filenames, line)
		}
	}
}
