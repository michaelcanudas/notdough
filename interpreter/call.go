package interpreter

import (
	"errors"
	"fmt"
	"michaelcanudas.dough/ast"
)

type Call struct {
	Identifier string
}

func init() {
	RegisterInstruction("call", func(arg ast.Node) Instruction {
		var identifier string

		if arg == nil {
			identifier = ""
		} else {
			identifier = arg.(ast.IlVariableNode).VarName.Value
		}

		return Call{
			Identifier: identifier,
		}
	})
}

func (c Call) execute(ctx *Context) {
	ctx.LocalStack.Push(make(map[string]int64))
	ctx.ReturnStack.Push(ctx.Pointer)

	if c.Identifier == "" {
		addr := ctx.ExecutionStack.Pop()

		// make sure it's a func we're calling
		_, ok := ctx.Instructions[addr].(Func)

		if !ok {
			panic("Indirect call tried to call a non-function address!")
		}

		ctx.Pointer = addr
	} else {
		val, ok := ctx.FuncMap[c.Identifier]
		if !ok {
			panic(errors.New(fmt.Sprintf("unknown function %s", c.Identifier)))
		}

		ctx.Pointer = val
	}
}
