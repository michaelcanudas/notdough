package parser

import "michaelcanudas.dough/ast"

func typ() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		return either(func(input []string) (ast.Node, []string, bool) {
			parsers := []Parser[ast.Node] {
				keyword("int"),
			}
			
			value, rest, ok := either(parsers...)(input)
			if !ok {
				return nil, input, ok
			}
			
			return ast.TypeNode{
				Value: value.(ast.KeywordNode),
				}, rest, ok
		}, additive())(input)
	}
}
