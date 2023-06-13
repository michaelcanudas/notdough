package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func block(node ast.BlockNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction
	
	for _, v := range node.Value {
		instructions = append(instructions, Compile(v)...)
	}
	
	return instructions
}
