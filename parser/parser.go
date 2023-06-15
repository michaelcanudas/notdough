package parser

import (
	"michaelcanudas.dough/ast"
)

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

func IlParse(input []string) []ast.IlInstructionNode {
	var (
		node ast.Node
		rest = input
		ok = true
	)
	var nodes []ast.IlInstructionNode
	for len(rest) > 0 {
		node, rest, ok = instruction()(rest)

		if !ok {
			return nil
		}

		ilnode, ok := node.(ast.IlInstructionNode)

		if !ok {
			return nil
		}

		nodes = append(nodes, ilnode)
	}

	return nodes
}