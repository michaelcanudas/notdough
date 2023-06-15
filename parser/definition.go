package parser

import "michaelcanudas.dough/ast"

func definition() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		nodes, rest, ok := sequence(
			keyword("let"),
			identifier(),
			optional(annotation()),
			symbol("="),
			expression())(input)
		if !ok {
			return nil, input, ok
		}

		if nodes[2] == nil {
			nodes[2] = ast.TypeNode{}
		}

		return ast.DefinitionNode{
			Keyword:    nodes[0].(ast.KeywordNode),
			Identifier: nodes[1].(ast.IdentifierNode),
			Type:       nodes[2].(ast.TypeNode),
			Assignment: nodes[3].(ast.SymbolNode),
			Value:      nodes[4],
		}, rest, ok
	}, ilExpression())
}
