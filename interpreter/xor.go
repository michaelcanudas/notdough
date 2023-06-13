package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Xor struct {
	Bitwise bool
}

func init() {
	RegisterInstruction("xor", func(arg ast.Node) Instruction {
		return Xor{
			Bitwise: HasBitwiseModifier(arg),
		}
	})
}

func (x Xor) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	var value int64
	if x.Bitwise {
		value = x.executeBitwise(a, b)
	} else {
		value = x.executeLogical(a, b)
	}

	ctx.ExecutionStack.Push(value)
}

func (x Xor) executeBitwise(a int64, b int64) int64 {
	return a ^ b
}

func (x Xor) executeLogical(a int64, b int64) int64 {
	if a != 0 {
		a = 1
	}
	if b != 0 {
		b = 1
	}

	result := a != b

	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	return value
}
