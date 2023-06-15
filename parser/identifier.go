package parser

import (
	"regexp"

	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/utils"
)

func identifier() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		if utils.Contains(keywords(), input[0]) {
			return nil, input, false
		}

		match, _ := regexp.MatchString("^[_a-zA-Z][_a-zA-Z0-9]*$", input[0])
		if !match {
			return nil, input, match
		}

		return ast.IdentifierNode{
			Value: input[0],
		}, input[1:], match
	}, number())
}
