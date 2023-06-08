package interpreter

type If struct {
}

func (i If) execute(ctx *Context) {
	// if instructions pop a value and skip to their corresponding 'endif' if it was 0
	end := FindEndif(ctx)
	if ctx.ExecutionStack.Pop() == 0 {
		ctx.Pointer = end
	}
}
