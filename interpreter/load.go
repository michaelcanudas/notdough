package interpreter

import (
	"errors"
	"fmt"
)

type Load struct {
	Identifier string
}

func init() {
	RegisterInstruction("load", func(fields []string) Instruction {
		return Load{
			Identifier: fields[1],
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
