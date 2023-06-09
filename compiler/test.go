package main

import (
	"fmt"
	
	"michaelcanudas.dough/compiler/parser"
)

func main() {
	program := parser.Parse([]string{
		"let",
		"main",
		"=",
		"{",
		"}",
	})
	
	fmt.Println(program)
}