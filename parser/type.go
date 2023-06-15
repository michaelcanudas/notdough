package parser

import "michaelcanudas.dough/ast"

func typ() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		term, rest, ok := either(keyword("int"))(input)
		if !ok {
			return nil, input, ok
		}

		return typHelper(ast.KeywordTypeNode{
			Type: term.(ast.KeywordNode),
		})(rest)
	}, additive())
}

func typHelper(term ast.Node) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		_, rest, ok := symbol("(")(input)
		if !ok {
			return term, input, true
		}

		ts, rest, ok := many(typ())(rest)
		if !ok {
			return term, input, true
		}

		_, rest, ok = symbol(")")(rest)
		if !ok {
			return term, input, true
		}

		return typHelper(ast.FunctionTypeNode{
			Type:     term,
			ArgTypes: argTypes,
		})(rest)
	}
}
