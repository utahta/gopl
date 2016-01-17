package main

import (
	"./testexp"
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func EmployeeByID(id int) *Employee {
	a := Employee{}
	a.ID = id
	a.Name = "test"
	a.Salary = 100
	return &a
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	//X int
}

func main() {
	EmployeeByID(1).Salary = 10
	fmt.Println(EmployeeByID(1).Salary)
	a := testexp.Crew{Name: "name1"}
	a.Name = "name2"
	//a.age = 20 // compile error
	fmt.Println(a)

	b := Wheel{}
	b.X = 10
	fmt.Printf("%#v\n", b)
}
