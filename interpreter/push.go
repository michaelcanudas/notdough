package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Push struct {
	Value int64
}

func init() {
	RegisterInstruction("push", func(arg ast.Node) Instruction {
		return Push{
			Value: arg.(ast.NumberNode).Value,
		}
	})
}

func (p Push) execute(ctx *Context) {
	ctx.ExecutionStack.Push(p.Value)
}
