package interpreter

type Push struct {
	Value int64
}

func (p Push) execute(ctx *Context) {
	ctx.Stack = append(ctx.Stack, p.Value)
}
