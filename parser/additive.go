package parser

import (
	"fmt"
	"michaelcanudas.dough/ast"
)

func additive() Parser[ast.Node] {
	var a Parser[ast.Node]
	a = either(func(input []string) (ast.Node, []string, bool) {
		term, rest, ok := multiplicative()(input)
		if !ok {
			return nil, input, ok
		}
		return additiveHelper(term)(rest)
	}, multiplicative())

	return a
}

// recursively parses operands til we can't no more
func additiveHelper(term ast.Node) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		t, rest, ok := either(add(term), sub(term))(input)

		if !ok {
			return term, input, true
		}

		return additiveHelper(t)(rest)
	}
}

func add(lhs ast.Node) Parser[ast.Node] {
	return binary("+", multiplicative())(lhs)
}

func sub(lhs ast.Node) Parser[ast.Node] {
	return binary("-", multiplicative())(lhs)
}
