package interpreter

type Div struct {
}

func init() {
	RegisterInstruction("div", func(fields []string) Instruction {
		return Div{}
	})
}

func (d Div) execute(ctx *Context) {
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()
	ctx.ExecutionStack.Push(a / b)
}
