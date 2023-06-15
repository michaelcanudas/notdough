package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func definition(node ast.DefinitionNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	switch node.Type.(type) {
	case ast.KeywordTypeNode:
		instructions = append(instructions, Compile(node.Value)...)
		instructions = append(instructions, interpreter.Store{
			Identifier: node.Identifier.Value,
		})
	case ast.FunctionTypeNode:
		instructions = append(instructions, interpreter.Func{
			Identifier: node.Identifier.Value,
		})
		instructions = append(instructions, Compile(node.Value)...)
		instructions = append(instructions, interpreter.Endfunc{})
	}

	return instructions
}
