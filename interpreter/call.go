package interpreter

type Call struct {
	Identifier string
}

func (c Call) execute(ctx *Context) {
	ctx.ReturnStack.Push(ctx.Pointer)
	ctx.Pointer = ctx.FuncMap[c.Identifier]
}
