package interpreter

type Cont struct {
}

func (c Cont) execute(ctx *Context) {
	loop := FindLoop(ctx)
	ctx.Pointer = loop
}
