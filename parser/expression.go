package parser

import "michaelcanudas.dough/ast"

func expression() Parser[ast.Node] {
	return func(input[] string) (ast.Node, []string, bool) {
		return block()(input)
	}
}
