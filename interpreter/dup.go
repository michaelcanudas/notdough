package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Dup struct {
}

func init() {
	RegisterInstruction("dup", func(arg ast.Node) Instruction {
		return Dup{}
	})
}

func (d Dup) execute(ctx *Context) {
	var value = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(value)
	ctx.ExecutionStack.Push(value)
}
