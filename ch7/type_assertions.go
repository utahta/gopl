package main

import "fmt"

type Aaa interface {
	Hoge() string
}

type Aaa2 interface {
	Hoge() string
	Foo() string
}

type Bbb struct {}

func (b Bbb) Hoge() string {
	return "bbb"
}

func (b Bbb) Foo() string {
	return "bbb foo"
}

type Ccc struct {}

func (c Ccc) Hoge() string {
	return "ccc"
}

func main() {
	var a Aaa
	a = Bbb{}
	fmt.Println("a", a.Hoge())

	var b Bbb
	b = a.(Bbb)
	fmt.Println("b", b.Hoge())

	// panic
	//var e Ccc
	//e = a.(Ccc)
	//fmt.Println(e.Hoge())

	var a2 Aaa2
	a2 = a.(Aaa2)
	fmt.Println("a2", a2.Hoge())

	// compile error
	//var e Aaa2
	//e = a.(Ccc)
	//fmt.Println(e.Hoge())

	// not panic
	var c Ccc
	c, ok := a.(Ccc)
	fmt.Println("c", c, ok)
}
