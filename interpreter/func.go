package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Func struct {
	Identifier string
}

func init() {
	RegisterInstruction("func", func(arg ast.Node) Instruction {
		return Func{
			Identifier: arg.(ast.IlVariableNode).VarName.Value,
		}
	})
}

func (f Func) execute(ctx *Context) {
	// functions, when executed, just skip to their corresponding 'endfunc'
	// instruction (so that they are not executed until they are called)
	ctx.Pointer = FindEndfunc(ctx)
}

func (f Func) getName() string {
	return f.Identifier
}