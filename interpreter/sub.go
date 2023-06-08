package interpreter

type Sub struct {
}

func init() {
	RegisterInstruction("sub", func(fields []string) Instruction {
		return Sub{}
	})
}

func (s Sub) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a - b)
}
