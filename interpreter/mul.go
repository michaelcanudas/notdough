package interpreter

type Mul struct {
}

func (m Mul) execute(ctx *Context) {
	ctx.Stack = append(ctx.Stack[:len(ctx.Stack) - 2], ctx.Stack[len(ctx.Stack) - 1] * ctx.Stack[len(ctx.Stack) - 2])
}
