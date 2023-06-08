package interpreter

type And struct {
}

func (and And) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	result := (a != 0) && (b != 0)

	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	ctx.ExecutionStack.Push(value)
}
