package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Sub struct {
}

func init() {
	RegisterInstruction("sub", func(arg ast.Node) Instruction {
		return Sub{}
	})
}

func (s Sub) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a - b)
}
