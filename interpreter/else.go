package interpreter

import "michaelcanudas.dough/ast"

type Else struct {
}

func init() {
	RegisterInstruction("else", func(node ast.Node) Instruction {
		return Else{}
	})
}

func (e Else) execute(ctx *Context) {
	ctx.Pointer = FindElseOrEndif(ctx)
}
