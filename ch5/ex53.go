package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visit(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	} else if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}
