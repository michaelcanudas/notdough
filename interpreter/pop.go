package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Pop struct {
	Count int64
}

func init() {
	RegisterInstruction("pop", func(arg ast.Node) Instruction {
		return Pop{
			Count: arg.(ast.NumberNode).Value,
		}
	})
}

func (p Pop) execute(ctx *Context) {
	ctx.ExecutionStack.Pop()
}
