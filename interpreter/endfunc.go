package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Endfunc struct {
}

func init() {
	RegisterInstruction("endfunc", func(arg ast.Node) Instruction {
		return Endfunc{}
	})
}

func (e Endfunc) execute(ctx *Context) {
	// 'endfunc' has the same behavior as 'Ret'
	ret := Ret{}
	ret.execute(ctx)
}
