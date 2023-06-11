package interpreter

type Inc struct {
}

func init() {
	RegisterInstruction("inc", func(fields []string) Instruction {
		return Inc{}
	})
}

func (i Inc) execute(ctx *Context) {
	ctx.ExecutionStack.Push(ctx.ExecutionStack.Pop() + 1)
}
