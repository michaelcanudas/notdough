package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func identifier(node ast.IdentifierNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, interpreter.Load{
		Identifier: node.Value,
	})
	
	return instructions
}
