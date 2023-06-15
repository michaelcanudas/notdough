package parser

import "michaelcanudas.dough/ast"

func Attribute(argParser Parser[ast.Node]) Parser[ast.Node] {
	return convert(
		sequence[ast.Node](
			symbol("@"),
			identifier(),
			argParser,
		),
		func(nodes []ast.Node) ast.Node {
			return ast.AttributeNode{
				Name:  nodes[1].(ast.IdentifierNode),
				Value: nodes[2],
			}
		},
	)
}

func IlAttribute() Parser[ast.Node] {
	return Attribute(
		optional(
			number(),
		),
	)
}

func CodeAttribute() Parser[ast.Node] {
	return Attribute(
		optional(
			either(
				number(),
				typ(),
			),
		),
	)
}

func convert[From any, To any](parser Parser[From], mapper func(From) To) Parser[To] {
	return func(input []string) (To, []string, bool) {
		node, rest, ok := parser(input)

		if !ok {
			var t To
			return t, input, false
		}

		return mapper(node), rest, true
	}
}
