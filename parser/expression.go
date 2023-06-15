package parser

import "michaelcanudas.dough/ast"

func expression() Parser[ast.Node] {
	return block()
}
