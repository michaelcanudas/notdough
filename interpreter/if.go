package interpreter

import (
	"michaelcanudas.dough/ast"
)

type If struct {
}

func init() {
	RegisterInstruction("if", func(arg ast.Node) Instruction {
		return If{}
	})
}

func (i If) execute(ctx *Context) {
	// if instructions pop a value and skip to their corresponding 'endif' if it was 0
	end := FindElseOrEndif(ctx)
	if ctx.ExecutionStack.Pop() == 0 {
		ctx.Pointer = end
	}
}
