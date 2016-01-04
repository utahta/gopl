package main

import (
	"os"
	"strings"
	"testing"
)

func Echo1() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}
