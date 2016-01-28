package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, cource := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, cource)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(mm map[string][]string)

	visitAll = func(mm map[string][]string) {
		for key, items := range mm {
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					if _, ok := m[item]; ok {
						mmm := map[string][]string{item: m[item]}
						visitAll(mmm)
					}
					order = append(order, item)
				}
			}

			if !seen[key] {
				seen[key] = true
				order = append(order, key)
			}
		}
	}

	visitAll(m)
	return order
}
