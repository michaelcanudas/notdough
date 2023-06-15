package parser

import "michaelcanudas.dough/ast"

func IfExpr() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		terms, rest, ok := sequence(
			keyword("if"),
			expression(),
			expression(),
			optional(
				convert(
					sequence(
						keyword("else"),
						expression(),
					),
					func(nodes []ast.Node) ast.Node {
						return nodes[1]
					},
				),
			),
		)(input)

		if !ok {
			return nil, input, false
		}

		var elseBlock ast.Node = nil

		if len(terms) >= 4 {
			elseBlock = terms[3]
		}

		node := ast.If{
			Condition: terms[1],
			Body:      terms[2],
			ElseBlock: elseBlock,
		}

		return node, rest, ok
	}, ilExpression())
}
