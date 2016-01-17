package main

import (
	"fmt"
)

func main() {
	a := map[string]int{
		"aaa": 20,
		"bbb": 30,
		"ccc": 40,
		"ddd": 50,
	}
	for name, age := range a {
		fmt.Println(name, age)
	}

	var b map[string]int
	//b["aaa"] = 10 // panic! map is nil
	b = map[string]int{}
	b["aaa"] = 10
	fmt.Println(b)

	c := map[string]map[string]bool{}
	fmt.Println(c["aaa"]["bbb"])
}
