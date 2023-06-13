package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func ret(node ast.ReturnNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Value)...)
	instructions = append(instructions, interpreter.Ret{})
	
	return instructions
}
