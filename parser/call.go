package parser

import (
	"michaelcanudas.dough/ast"
)

func call() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		term, rest, ok := parenthesis()(input)
		if !ok {
			return nil, input, ok
		}
		return callHelper(term)(rest)
	}, parenthesis())
}

func callHelper(term ast.Node) Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		params, rest, ok := clown()(input)
		if !ok {
			return term, input, true
		}

		return callHelper(ast.CallNode{
			Value:  term,
			Params: params,
		})(rest)
	}
}
