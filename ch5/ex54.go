package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func extractValid(n *html.Node, cond map[string]string) bool {
	if cond == nil {
		return true
	}
	for _, a := range n.Attr {
		if _, ok := cond[a.Key]; !ok {
			continue
		}
		if cond[a.Key] == a.Val {
			return true
		}
	}
	return false
}

func extractLink(n *html.Node, name string, attr string, cond map[string]string) {
	if n.Type == html.ElementNode && n.Data == name {
		if !extractValid(n, cond) {
			return
		}

		for _, a := range n.Attr {
			if a.Key == attr {
				fmt.Println(name, attr, a.Val)
			}
		}
	}
}

func visit(n *html.Node) {
	extractLink(n, "a", "href", nil)
	extractLink(n, "img", "src", nil)
	extractLink(n, "script", "src", nil)
	extractLink(n, "link", "href", map[string]string{"rel": "stylesheet"})

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
