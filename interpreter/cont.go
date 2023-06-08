package interpreter

type Cont struct {
}

func (c Cont) execute(ctx *Context) {
	// continue instruction skips to the next iteration of the current loop

	// find the loop instruction
	loop := FindLoop(ctx)
	// and go to it
	ctx.Pointer = loop
}
