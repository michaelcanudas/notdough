package interpreter

type Not struct {
}

func (n Not) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()

	result := a != 0

	var value int64
	if result {
		value = 0
	} else {
		value = 1
	}

	ctx.ExecutionStack.Push(value)
}
