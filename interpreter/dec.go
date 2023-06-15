package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Dec struct {
}

func init() {
	RegisterInstruction("dec", func(arg ast.Node) Instruction {
		return Dec{}
	})
}

func (d Dec) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() - 1)
}
