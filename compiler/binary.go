package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func binary(node ast.BinaryNode) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(node.Rhs)...)
	instructions = append(instructions, Compile(node.Lhs)...)
	switch node.Op.Value {
	case "+":
		instructions = append(instructions, interpreter.Add{})
	case "-":
		instructions = append(instructions, interpreter.Sub{})
	case "*":
		instructions = append(instructions, interpreter.Mul{})
	case "/":
		instructions = append(instructions, interpreter.Div{})
	}
	
	return instructions
}
