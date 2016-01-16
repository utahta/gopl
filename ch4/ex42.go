package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var useType string
	flag.StringVar(&useType, "sha", "256", "256, 384, 512. ")
	flag.Parse()

	arg := []byte(flag.Arg(0))
	switch useType {
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(arg))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(arg))
	default:
		fmt.Printf("%x\n", sha256.Sum256(arg))
	}
}
