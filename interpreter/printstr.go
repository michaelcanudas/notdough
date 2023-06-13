package interpreter

import (
	"michaelcanudas.dough/ast"
)

func init() {
	RegisterInstruction("printstr", func(arg ast.Node) Instruction {
		return Printstr{}
	})
}

type Printstr struct {
}

func (p Printstr) execute(ctx *Context) {
	_ = ctx.ExecutionStack.Pop()
}