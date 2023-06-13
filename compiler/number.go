package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func number(node ast.NumberNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, interpreter.Push{
		Value: node.Value,
	})
	
	return instructions
}
