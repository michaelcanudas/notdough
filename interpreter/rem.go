package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Rem struct {
}

func init() {
	RegisterInstruction("rem", func(arg ast.Node) Instruction {
		return Rem{}
	})
}

func (r Rem) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a % b)
}
