package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func ifExpr(i ast.If) []interpreter.Instruction {
	var instructions []interpreter.Instruction

	instructions = append(instructions, Compile(i.Condition)...)
	instructions = append(instructions, interpreter.If{})
	instructions = append(instructions, Compile(i.Body)...)

	if i.ElseBlock != nil {
		instructions = append(instructions, interpreter.Else{})
		instructions = append(instructions, Compile(i.ElseBlock)...)
	}

	instructions = append(instructions, interpreter.Endif{})
	return instructions
}
