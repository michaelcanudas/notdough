package interpreter

import (
	"fmt"
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
	strLocation := ctx.ExecutionStack.Pop()
	str := ctx.Instructions[strLocation+1].(Text)
	fmt.Print(str.Value)
}
