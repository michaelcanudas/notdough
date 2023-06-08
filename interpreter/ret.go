package interpreter

type Ret struct {
}

func (r Ret) execute(ctx *Context) {
	ctx.Pointer = ctx.ReturnStack[len(ctx.ReturnStack) - 1]
	ctx.ReturnStack = ctx.ReturnStack[:len(ctx.ReturnStack) - 1]
}
