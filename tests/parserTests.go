package tests

import (
	"fmt"
	"michaelcanudas.dough/parser"
)

func init() {
	RegisterTest("parser.main", func() {
		input := []string{
			"let",
			"x",
		}
		node, rest, ok := parser.Definition()(input)

		fmt.Println(node)
		fmt.Println(rest)
		fmt.Println(ok)
	})
}
