package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func el(y ...int) {
	fmt.Println(y)
}

func main() {
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	b := a[1:2]
	b[0] = "B"
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))

	//c := b[5] // outof range
	//fmt.Println(c)
	c := b[1:5]
	fmt.Println(c, len(c), cap(c))

	d1 := a
	d2 := a[:]
	a[0] = "A"
	fmt.Println(d1[0] == d2[0])

	e := "abcdefg"
	e2 := e[1:5]
	fmt.Printf("%s, %T\n", e2, e2)

	f := []int{0, 1, 2, 3, 4, 5}
	reverse(f)
	fmt.Println(f)

	g := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(g) // compile error
	reverse(g[:])
	fmt.Println("g:", g)

	f = []int{0, 1, 2, 3, 4, 5}
	reverse(f[:2])
	fmt.Println(f)
	reverse(f[2:])
	fmt.Println(f)
	reverse(f)
	fmt.Println(f)

	reverse(nil)

	h1 := []int{0, 1, 2, 3}
	h2 := make([]int, 4)
	copy(h2, h1)
	fmt.Println(h2)
	copy(h1, h1)
	fmt.Println(h1)
	h3 := []int{10, 20}
	copy(h3, h1)
	fmt.Println(h3)
	h2 = append(h2, 4)
	fmt.Println(h2)

	i1 := [...]int{0, 1, 2, 3}
	i2 := i1[:]
	i2[0] = 10
	fmt.Println(i1, "views")
	i2 = append(i2, 4)
	i2[0] = 100
	fmt.Println(i1, "not views")

	el(1, 2, 3)
}
