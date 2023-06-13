package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func print(node ast.PrintNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Value)...)
	instructions = append(instructions, interpreter.Print{})
	
	return instructions
}
