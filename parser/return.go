package parser

import (
	"michaelcanudas.dough/ast"
)

func ret() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			parsers := []Parser[ast.Node] {
				keyword("return"),
				expression(),
			}
			
			nodes, rest, ok := sequence(parsers...)(input)
			if !ok {
				return nil, input, ok
			}
			
			return ast.ReturnNode{
				Keyword: nodes[0].(ast.KeywordNode),
				Value: nodes[1],
			}, rest, ok
		}, typ())(input)
	}
}
