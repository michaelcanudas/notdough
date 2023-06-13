package parser

import (
	"michaelcanudas.dough/ast"
)

func symbol(symbol string) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		value, rest, ok := str(symbol)(input)
		if !ok {
			return nil, input, ok
		}

		return ast.SymbolNode{
			Value: value,
			}, rest, ok
	}
}