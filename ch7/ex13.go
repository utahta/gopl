package main

import (
	"fmt"
	"github.com/utahta/gopl/ch7/eval"
)

func main() {
	expr, _ := eval.Parse("pow(x + pow(y))")
	fmt.Println(expr.String())
}
