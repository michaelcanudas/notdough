package parser

import "michaelcanudas.dough/ast"

func block() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		_, rest, ok := optional(typ())(input)

		open, rest, ok := str("{")(rest)
		if !ok {
			return nil, input, ok
		}

		value, rest, ok := many(expression())(rest)
		if !ok {
			return nil, input, ok
		}

		close, rest, ok := str("}")(rest)
		if !ok {
			return nil, input, ok
		}

		return ast.BlockNode{
			Open:  open,
			Value: value,
			Close: close,
		}, rest, ok
	}, definition())
}
