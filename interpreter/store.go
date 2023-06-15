package interpreter

import (
	"michaelcanudas.dough/ast"
)

type Store struct {
	Identifier string
}

func init() {
	RegisterInstruction("store", func(arg ast.Node) Instruction {
		return Store{
			Identifier: arg.(ast.IlVariableNode).VarName.Value,
		}
	})
}

func (s Store) execute(ctx *Context) {
	ctx.LocalStack.Peek()[s.Identifier] = ctx.ExecutionStack.Pop()
}
