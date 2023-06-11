package interpreter

type Ret struct {
}

func init() {
	RegisterInstruction("ret", func(fields []string) Instruction {
		return Ret{}
	})
}

func (r Ret) execute(ctx *Context) {
	// ret simply jumps to the instruction at the top of the return stack
	ctx.LocalStack.Pop()
	ctx.Pointer = ctx.ReturnStack.Pop()
}
