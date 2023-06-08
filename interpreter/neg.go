package interpreter

type Neg struct {
}

func init() {
	RegisterInstruction("neg", func(fields []string) Instruction {
		return Neg{}
	})
}

func (n Neg) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(-a)
}
