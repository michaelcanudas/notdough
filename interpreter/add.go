package interpreter

type Add struct {
}

func (a Add) execute(ctx *Context) {
	ctx.Stack = append(ctx.Stack[:len(ctx.Stack) - 2], ctx.Stack[len(ctx.Stack) - 1] + ctx.Stack[len(ctx.Stack) - 2])
}
