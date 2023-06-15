package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func call(node ast.CallNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Value)...)
	instructions = append(instructions, interpreter.Call{})

	return instructions
}
