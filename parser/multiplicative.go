package parser

import "michaelcanudas.dough/ast"

func multiplicative() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		factor, rest, ok := negative()(input)
		if !ok {
			return nil, input, ok
		}

		return multiplicativeHelper(factor)(rest)
	}, negative())
}

func multiplicativeHelper(factor ast.Node) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		t, rest, ok := either(mul(factor), div(factor))(input)

		if !ok {
			return factor, input, true
		}

		return multiplicativeHelper(t)(rest)
	}
}

func mul(lhs ast.Node) Parser[ast.Node] {
	return binary("*", negative())(lhs)
}

func div(lhs ast.Node) Parser[ast.Node] {
	return binary("/", negative())(lhs)
}
