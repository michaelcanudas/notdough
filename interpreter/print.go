package interpreter

import "fmt"

type Print struct {
}

func (p Print) execute(ctx *Context) {
	fmt.Println(ctx.ExecutionStack.Pop())
}
