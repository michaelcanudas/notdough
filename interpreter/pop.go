package interpreter

type Pop struct {
	Count int64
}

func (p Pop) execute(ctx *Context) {
	ctx.ExecutionStack.Pop()
}
