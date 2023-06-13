package parser

import "michaelcanudas.dough/ast"

type Parser[T any] func([]string) (T, []string, bool)

func Parse(input []string) ast.Node {
	node, rest, ok := expression()(input)

	if !ok {
		return nil
	}

	if len(rest) != 0 {
		return nil
	}

	return node
}
