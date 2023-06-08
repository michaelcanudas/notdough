package interpreter

import "fmt"

type Print struct {
}

func (p Print) execute(ctx *Context) {
	fmt.Println(ctx.Stack[len(ctx.Stack) - 1])
	ctx.Stack = ctx.Stack[:len(ctx.Stack) - 1]
}
