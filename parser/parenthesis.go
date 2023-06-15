package parser

import "michaelcanudas.dough/ast"

func parenthesis() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		nodes, rest, ok := sequence(
			symbol("("),
			expression(),
			symbol(")"))(input)
		if !ok {
			return nil, input, ok
		}

		return nodes[1], rest, ok
	}, identifier())
}
