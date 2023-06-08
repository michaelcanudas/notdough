package interpreter

type Dec struct {
}

func (d Dec) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() - 1)
}
