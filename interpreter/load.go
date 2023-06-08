package interpreter

import "errors"

type Load struct {
	Identifier string
}

func (l Load) execute(ctx *Context) {
	val, ok := ctx.LocalStack[len(ctx.LocalStack)-1][l.Identifier]
	if !ok {
		panic(errors.New("unknown identifier"))
	}

	ctx.ExecutionStack.Push(val)
}
