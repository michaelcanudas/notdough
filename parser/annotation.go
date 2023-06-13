package parser

import "michaelcanudas.dough/ast"

func annotation() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		parsers := []Parser[ast.Node] {
			symbol(":"),
			typ(),
		}
		
		nodes, rest, ok := sequence(parsers...)(input)
		if !ok {
			return nil, input, ok
		}
		
		return nodes[1], rest, ok
	}
}
