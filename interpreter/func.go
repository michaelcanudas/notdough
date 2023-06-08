package interpreter

type Func struct {
	Identifier string
}

func (f Func) execute(ctx *Context) {
	// functions, when executed, just skip to their corresponding 'endfunc'
	// instruction (so that they are not executed until they are called)
	ctx.Pointer = FindEndfunc(ctx)
}
