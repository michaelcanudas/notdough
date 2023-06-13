package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Endif struct {
}

func init() {
	RegisterInstruction("endif", func(arg ast.Node) Instruction {
		return Endif{}
	})
}

func (e Endif) execute(ctx *Context) {
	// nop
}
