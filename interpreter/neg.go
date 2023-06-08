package interpreter

type Neg struct {
}

func (n Neg) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(-a)
}
