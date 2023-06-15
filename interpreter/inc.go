package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Inc struct {
}

func init() {
	RegisterInstruction("inc", func(arg ast.Node) Instruction {
		return Inc{}
	})
}

func (i Inc) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() + 1)
}
