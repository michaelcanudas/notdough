package interpreter

type If struct {
}

func (i If) execute(ctx *Context) {
	end := FindEndif(ctx)
	if ctx.Stack[len(ctx.Stack) - 1] == 0 {
		ctx.Pointer = end
	}
		
	ctx.Stack = ctx.Stack[:len(ctx.Stack) - 1]
}
