package parser

import (
	"michaelcanudas.dough/ast"
)

func keyword(keyword string) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		value, rest, ok := str(keyword)(input)
		if !ok {
			return nil, input, ok
		}

		return ast.KeywordNode{
			Value: value,
		}, rest, ok
	}
}
