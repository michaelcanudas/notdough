package interpreter

type Dup struct {
}

func init() {
	RegisterInstruction("dup", func(fields []string) Instruction {
		return Dup{}
	})
}

func (d Dup) execute(ctx *Context) {
	var value = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(value)
	ctx.ExecutionStack.Push(value)
}
