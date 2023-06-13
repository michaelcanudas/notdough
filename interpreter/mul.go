package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Mul struct {
}

func init() {
	RegisterInstruction("mul", func(arg ast.Node) Instruction {
		return Mul{}
	})
}

func (m Mul) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a * b)
}
