package parser

import "michaelcanudas.dough/ast"

func parenthesis() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			parsers := []Parser[ast.Node] {
				symbol("("),
				expression(),
				symbol(")"),
			}
			
			nodes, rest, ok := sequence(parsers...)(input)
			if !ok {
				return nil, input, ok
			}
			
			return nodes[1], rest, ok
		}, number())(input)
	}
}
