package interpreter

type Not struct {
	Bitwise bool
}

func (n Not) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()

	var value int64
	if n.Bitwise {
		value = n.executeBitwise(a)
	} else {
		value = n.executeLogical(a)
	}

	ctx.ExecutionStack.Push(value)
}

func (n Not) executeBitwise(a int64) int64 {
	return ^a
}

func (n Not) executeLogical(a int64) int64 {
	result := a != 0

	var value int64
	if result {
		value = 0
	} else {
		value = 1
	}

	return value
}
