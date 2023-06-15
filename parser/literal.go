package parser

import (
	"michaelcanudas.dough/ast"
	"strings"
)

func literal() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		if len(input) == 0 {
			return nil, input, false
		}

		if strings.HasPrefix(input[0], `"`) && strings.HasSuffix(input[0], `"`) {
			result := ast.LiteralNode{
				Value: input[0][1:len(input[0])-1],
			}
			
			return result, input[1:], true
		}
		
		return nil, input, false
	}
}