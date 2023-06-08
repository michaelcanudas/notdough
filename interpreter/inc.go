package interpreter

type Inc struct {
}

func (i Inc) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() + 1)
}
