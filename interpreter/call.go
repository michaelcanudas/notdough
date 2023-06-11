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
		var identifier string

		if len(fields) > 1 {
			identifier = fields[1]
		} else {
			identifier = ""
		}

		return Call{
			Identifier: identifier,
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
