package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Endloop struct {
}

func init() {
	RegisterInstruction("endloop", func(arg ast.Node) Instruction {
		return Endloop{}
	})
}

func (e Endloop) execute(ctx *Context) {
	// nop
}
