package parser

import "michaelcanudas.dough/ast"

func block() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			open, rest, ok := str("{")(input)
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

			return ast.BlockNode {
				Open: open,
				Value: value,
				Close: close,
			}, rest, ok
		}, definition())(input)
	}
}
