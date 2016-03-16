package main

import (
	"bufio"
	"os"
	"fmt"
	"github.com/utahta/gopl/ch7/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		expr, err := eval.Parse(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err = expr.Check(map[eval.Var]bool{}); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(expr.Eval(eval.Env{}))
	}
}
