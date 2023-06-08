package interpreter

import (
	"errors"
	"fmt"
)

type Call struct {
	Identifier string
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
