package parser

import "michaelcanudas.dough/ast"

func unary(operator string, next Parser[ast.Node]) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		nodes, rest, ok := sequence(
			symbol(operator),
			next)(input)
		if !ok {
			return nil, input, ok
		}

		return ast.UnaryNode{
			Op:    nodes[0].(ast.SymbolNode),
			Value: nodes[1],
		}, rest, ok
	}
}
