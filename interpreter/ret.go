package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Ret struct {
}

func init() {
	RegisterInstruction("ret", func(arg ast.Node) Instruction {
		return Ret{}
	})
}

func (r Ret) execute(ctx *Context) {
	// ret simply jumps to the instruction at the top of the return stack
	ctx.LocalStack.Pop()
	ctx.Pointer = ctx.ReturnStack.Pop()
}
