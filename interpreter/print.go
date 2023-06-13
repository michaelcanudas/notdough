package interpreter

import (
	"fmt"
	"michaelcanudas.dough/ast"
)

type Print struct {
}

func init() {
	RegisterInstruction("print", func(arg ast.Node) Instruction {
		return Print{}
	})
}

func (p Print) execute(ctx *Context) {
	fmt.Println(ctx.ExecutionStack.Pop())
}
