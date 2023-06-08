package interpreter

import "errors"

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
		panic(errors.New("unknown identifier"))
	}

	ctx.ExecutionStack.Push(val)
}
