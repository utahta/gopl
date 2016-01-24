package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func elementsCount(elem map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elem[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementsCount(elem, c)
	}
	return elem
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, v := range elementsCount(map[string]int{}, doc) {
		fmt.Println(k, v)
	}
}
