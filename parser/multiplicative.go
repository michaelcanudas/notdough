package parser

import "michaelcanudas.dough/ast"

func multiplicative() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			factor, rest, ok := parenthesis()(input)
			if !ok {
				return nil, input, ok
			}

			return either(mul(factor), div(factor))(rest)
		}, parenthesis())(input)
	}
}

func mul(lhs ast.Node) Parser[ast.Node] {
	return binary("*", parenthesis())(lhs)
}

func div(lhs ast.Node) Parser[ast.Node] {
	return binary("/", parenthesis())(lhs)
}
