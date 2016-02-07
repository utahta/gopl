package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p Point) DupTest() {
	fmt.Println("duplicate test")
}

type Path []Point

func (path *Path) IsNil() bool {
	return *path == nil
}

func (path Path) Size() int {
	return len(path)
}

//func (p *Point) X() {
//	fmt.Println("compile error")
//}

type PointDummy struct {
	dummy int
}

func (p PointDummy) DupTest() {
	fmt.Println("duplicate test in dummy")
}

type ColoredPoint struct {
	Point
	PointDummy
}

func main() {
	var p1 Path
	fmt.Println(p1.IsNil())

	p2 := Path{
		{1, 1},
		{2, 2},
	}
	fmt.Println(p2.Size())

	// compile error
	//c := ColoredPoint{}
	//c.DupTest()
}
