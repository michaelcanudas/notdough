package interpreter

type Xor struct {
}

func (x Xor) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	if a != 0 {
		a = 1
	}
	if b != 0 {
		b = 1
	}

	result := a != b

	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	ctx.ExecutionStack.Push(value)
}
