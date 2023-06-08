package interpreter

type Func struct {
	Identifier string
}

func (f Func) execute(ctx *Context) {
	ctx.Pointer = FindEndfunc(ctx)
}
