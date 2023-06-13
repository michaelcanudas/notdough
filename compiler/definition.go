package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func definition(node ast.DefinitionNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Value)...)
	instructions = append(instructions, interpreter.Store{
		Identifier: node.Identifier.Value,
	})
	
	return instructions
}
