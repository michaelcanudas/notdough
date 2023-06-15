package interpreter

import (
	"errors"
	"fmt"
	"michaelcanudas.dough/ast"
)

type Load struct {
	Identifier string
}

func init() {
	RegisterInstruction("load", func(arg ast.Node) Instruction {
		return Load{
			Identifier: arg.(ast.IlVariableNode).VarName.Value,
		}
	})
}

func (l Load) execute(ctx *Context) {
	val, ok := ctx.LocalStack.Peek()[l.Identifier]
	if !ok {
		val, ok = ctx.FuncMap[l.Identifier]

		if !ok {
			panic(errors.New(fmt.Sprintf("unknown identifier %v", l.Identifier)))
		}
	}

	ctx.ExecutionStack.Push(val)
}
