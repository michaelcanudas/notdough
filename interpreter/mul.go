package interpreter

type Mul struct {
}

func init() {
	RegisterInstruction("mul", func(fields []string) Instruction {
		return Mul{}
	})
}

func (m Mul) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a * b)
}
