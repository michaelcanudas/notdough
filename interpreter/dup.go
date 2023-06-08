package interpreter

type Dup struct {
}

func (d Dup) execute(ctx *Context) {
	var value = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(value)
	ctx.ExecutionStack.Push(value)
}
