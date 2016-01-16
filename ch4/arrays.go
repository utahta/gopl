package main

import (
	"fmt"
)

func argCopy(array [4]int) {
	array = [4]int{2, 2, 2, 2}
	fmt.Println(array)
}

func argRef(array *[4]int) {
	*array = [4]int{2, 2, 2, 2}
	fmt.Println(*array)
}

func main() {
	//t := [5]int{1, 2, 3, 4.2, 5}
	//fmt.Println(t)

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//fmt.Println(a[len(a)])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{10, 20, 30}
	fmt.Println(q)

	r := [3]int{10, 20}
	fmt.Println(r)

	r = [3]int{3, 2, 1}
	fmt.Println(r)

	q2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(q2)

	q3 := [...]int{99: -1}
	fmt.Println(len(q3))

	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{1, 3}
	fmt.Println(a1 == a2, a1 == a3)
	//a4 := [3]int{1, 2}
	//fmt.Println(a1 == a4)

	aa := [4]int{1, 1, 1, 1}
	argCopy(aa)
	fmt.Println(aa)

	argRef(&aa)
	fmt.Println(aa)
}
