package main

import (
	"fmt"
	"github.com/utahta/gopl/ch7/eval"
)

func main() {
	expr, _ := eval.Parse("x ? y")
	fmt.Println(expr.Eval(eval.Env{"x":10, "y":20}))
}
