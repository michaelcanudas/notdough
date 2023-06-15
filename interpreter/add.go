package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Add struct {
}

func init() {
	RegisterInstruction("add", func(node ast.Node) Instruction {
		return Add{}
	})
}
func (add Add) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a + b)
}
