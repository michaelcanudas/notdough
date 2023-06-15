package parser

import "michaelcanudas.dough/ast"

func binary(operator string, next Parser[ast.Node]) func(node ast.Node) Parser[ast.Node] {
	return func(lhs ast.Node) Parser[ast.Node] {
		return func(input []string) (ast.Node, []string, bool) {
			nodes, rest, ok := sequence(
				symbol(operator),
				next)(input)
			if !ok {
				return nil, input, ok
			}

			return ast.BinaryNode{
				Op:  nodes[0].(ast.SymbolNode),
				Lhs: lhs,
				Rhs: nodes[1],
			}, rest, ok
		}
	}
}
