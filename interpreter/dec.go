package interpreter

type Dec struct {
}

func init() {
	RegisterInstruction("dec", func(fields []string) Instruction {
		return Dec{}
	})
}

func (d Dec) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() - 1)
}
