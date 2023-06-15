package parser

import (
	"michaelcanudas.dough/ast"
)

func ret() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		nodes, rest, ok := sequence(
			keyword("return"),
			expression())(input)
		if !ok {
			return nil, input, ok
		}

		return ast.ReturnNode{
			Keyword: nodes[0].(ast.KeywordNode),
			Value:   nodes[1],
		}, rest, ok
	}, typ())
}
