package interpreter

import (
	"errors"
	"fmt"
)

type Call struct {
	Identifier string
}

func init() {
	RegisterInstruction("call", func(fields []string) Instruction {
		return Call{
			Identifier: fields[1],
		}
	})
	RegisterInstruction("", func(fields []string) Instruction {
		return Call{
			Identifier: fields[0],
		}
	})
}

func (c Call) execute(ctx *Context) {
	ctx.LocalStack.Push(make(map[string]int64))
	ctx.ReturnStack.Push(ctx.Pointer)

	val, ok := ctx.FuncMap[c.Identifier]
	if !ok {
		panic(errors.New(fmt.Sprintf("unknown function %s", c.Identifier)))
	}

	ctx.Pointer = val
}
