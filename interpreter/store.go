package interpreter

type Store struct {
	Identifier string
}

func (s Store) execute(ctx *Context) {
	ctx.LocalStack.Peek()[s.Identifier] = ctx.ExecutionStack.Pop()
}
