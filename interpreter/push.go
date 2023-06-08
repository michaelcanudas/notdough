package interpreter

type Push struct {
	Value int64
}

func (p Push) execute(ctx *Context) {
	ctx.ExecutionStack.Push(p.Value)
}
