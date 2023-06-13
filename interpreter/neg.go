package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Neg struct {
}

func init() {
	RegisterInstruction("neg", func(arg ast.Node) Instruction {
		return Neg{}
	})
}

func (n Neg) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(-a)
}
