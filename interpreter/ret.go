package interpreter

type Ret struct {
}

func (r Ret) execute(ctx *Context) {
	// ret simply jumps to the instruction at the top of the return stack
	ctx.Pointer = ctx.ReturnStack.Pop()
}
