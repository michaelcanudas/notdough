package interpreter

type And struct {
	Bitwise bool
}

func init() {
	RegisterInstruction("and", func(fields []string) Instruction {
		return And{
			Bitwise: HasBitwiseModifier(fields),
		}
	})
}

func (and And) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	var value int64
	if and.Bitwise {
		value = and.executeBitwise(a, b)
	} else {
		value = and.executeLogical(a, b)
	}

	ctx.ExecutionStack.Push(value)
}

func (and And) executeBitwise(a int64, b int64) int64 {
	return a & b
}

func (and And) executeLogical(a int64, b int64) int64 {
	result := (a != 0) && (b != 0)

	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	return value
}
