package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
	"michaelcanudas.dough/utils"
)

func call(node ast.CallNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	params := utils.Reverse(node.Params)
	for _, p := range params {
		instructions = append(instructions, Compile(p)...)
	}

	instructions = append(instructions, Compile(node.Value)...)
	instructions = append(instructions, interpreter.Call{})

	return instructions
}
