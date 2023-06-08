package interpreter

import "errors"

type Call struct {
	Identifier string
}

func (c Call) execute(ctx *Context) {
	ctx.LocalStack.Push(make(map[string]int64))
	ctx.ReturnStack.Push(ctx.Pointer)

	val, ok := ctx.FuncMap[c.Identifier]
	if !ok {
		panic(errors.New("unknown function"))
	}

	ctx.Pointer = val
}
