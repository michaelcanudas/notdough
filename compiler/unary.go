package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func unary(node ast.UnaryNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Value)...)
	switch node.Op.Value {
	case "-":
		instructions = append(instructions, interpreter.Neg{})
	}

	return instructions
}
