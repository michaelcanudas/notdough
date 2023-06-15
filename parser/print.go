package parser

import (
	"michaelcanudas.dough/ast"
)

func print() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		nodes, rest, ok := sequence(
			keyword("print"),
			expression())(input)
		if !ok {
			return nil, input, ok
		}

		return ast.PrintNode{
			Keyword: nodes[0].(ast.KeywordNode),
			Value:   nodes[1],
		}, rest, ok
	}, ret())
}
