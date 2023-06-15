package parser

import "michaelcanudas.dough/ast"

func clown() Parser[[]ast.Node] {
	return func(input []string) ([]ast.Node, []string, bool) {
		_, rest, ok := symbol("(")(input)
		if !ok {
			return nil, input, ok
		}

		value, rest, ok := many(expression())(rest)
		if !ok {
			return nil, input, ok
		}

		_, rest, ok = symbol(")")(rest)
		if !ok {
			return nil, input, ok
		}

		return value, rest, ok
	}
}
