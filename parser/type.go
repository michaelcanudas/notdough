package parser

import "michaelcanudas.dough/ast"

func typ() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		value, rest, ok := either(keyword("int"))(input)
		if !ok {
			return nil, input, ok
		}

		return ast.TypeNode{
			Value: value.(ast.KeywordNode),
		}, rest, ok
	}, additive())
}
