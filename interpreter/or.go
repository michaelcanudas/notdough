package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Or struct {
	Bitwise bool
}

func init() {
	RegisterInstruction("or", func(arg ast.Node) Instruction {
		return Or{
			Bitwise: HasBitwiseModifier(arg),
		}
	})
}

func (or Or) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	var value int64
	if or.Bitwise {
		value = or.executeBitwise(a, b)
	} else {
		value = or.executeLogical(a, b)
	}

	ctx.ExecutionStack.Push(value)
}

func (or Or) executeBitwise(a int64, b int64) int64 {
	return a | b
}

func (or Or) executeLogical(a int64, b int64) int64 {
	result := (a != 0) || (b != 0)

	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	return value
}
