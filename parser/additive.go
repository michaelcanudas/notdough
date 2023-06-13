package parser

import "michaelcanudas.dough/ast"

func additive() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			term, rest, ok := multiplicative()(input)
			if !ok {
				return nil, input, ok
			}

			return either(add(term), sub(term))(rest)
		}, multiplicative())(input)
	}
}

func add(lhs ast.Node) Parser[ast.Node] {
	return binary("+", multiplicative())(lhs)
}

func sub(lhs ast.Node) Parser[ast.Node] {
	return binary("-", multiplicative())(lhs)
}
