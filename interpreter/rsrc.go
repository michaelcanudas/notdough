package interpreter

import (
	"michaelcanudas.dough/ast"
)

func init() {
	RegisterInstruction("rsrc", func(arg ast.Node) Instruction {
		return Rsrc{
			Name: arg.(ast.IlVariableNode).VarName.Value,
		}
	})
}

type Rsrc struct {
	Name string
}

func (r Rsrc) execute(ctx *Context) {
	// nop
}

func (r Rsrc) getName() string {
	return r.Name
}
