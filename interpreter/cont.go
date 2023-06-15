package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Cont struct {
}

func init() {
	RegisterInstruction("cont", func(arg ast.Node) Instruction {
		return Cont{}
	})
}
func (c Cont) execute(ctx *Context) {
	// continue instruction skips to the next iteration of the current loop

	// find the loop instruction
	loop := FindLoop(ctx)
	// and go to it
	ctx.Pointer = loop
}
