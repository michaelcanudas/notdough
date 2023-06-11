package main

import (
	"fmt"
	
	"michaelcanudas.dough/parser"
)

func main() {
	/*result, rest, ok := parser.Expression()([]string {
		"(",
		"2",
		"+",
		"2",
		")",
		"*",
		"3",
	})*/
	
	result, rest, ok := parser.Definition()([]string {
		"let",
		"x3094",
	})
	
	fmt.Println(result)
	fmt.Println(rest)
	fmt.Println(ok)
}