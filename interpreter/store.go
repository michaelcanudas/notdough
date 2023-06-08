package interpreter

type Store struct {
	Identifier string
}

func (s Store) execute(ctx *Context) {
	ctx.LocalStack[len(ctx.LocalStack)-1][s.Identifier] = ctx.ExecutionStack.Pop()
}
