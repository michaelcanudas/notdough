package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Loop struct {
}

func init() {
	RegisterInstruction("loop", func(arg ast.Node) Instruction {
		return Loop{}
	})
}

func (l Loop) execute(ctx *Context) {
	// nop
	// loops only serve as a header for other instructions (such as cont) to jump to
}
