package parser

import (
	"strconv"

	"michaelcanudas.dough/ast"
)

func number() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		if len(input) == 0 {
			return ast.NumberNode{}, input, false
		}

		value, err := strconv.ParseInt(input[0], 10, 64)
		if err != nil {
			return nil, input, false
		}

		return ast.NumberNode{
			Value: value,
			}, input[1:], true
	}
}
