package parser

import (
	"michaelcanudas.dough/ast"
)

func negative() Parser[ast.Node] {
	return either(negativeHelper(), call())
}

// recursively parses operands til we can't no more
func negativeHelper() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		term, rest, ok := unary("-", negativeHelper())(input)

		if !ok {
			return call()(input)
		}

		return term, rest, ok
	}
}
