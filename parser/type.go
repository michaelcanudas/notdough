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
		ts, rest, ok := clown()(input)
		if !ok {
			return term, input, true
		}

		var argTypes []ast.TypeNode
		for _, t := range ts {
			switch t.(type) {
			case ast.KeywordTypeNode:
			case ast.FunctionTypeNode:
				argTypes = append(argTypes, t.(ast.TypeNode))
			default:
				return term, input, true
			}
		}

		return typHelper(ast.FunctionTypeNode{
			Type:     term,
			ArgTypes: argTypes,
		})(rest)
	}
}
