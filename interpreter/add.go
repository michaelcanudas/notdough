package interpreter

type Add struct {
}

func init() {
	RegisterInstruction("add", func(fields []string) Instruction {
		return Add{}
	})
}
func (add Add) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a + b)
}
