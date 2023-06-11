package interpreter

type Store struct {
	Identifier string
}

func init() {
	RegisterInstruction("store", func(fields []string) Instruction {
		return Store{
			Identifier: fields[1],
		}
	})
}

func (s Store) execute(ctx *Context) {
	ctx.LocalStack.Peek()[s.Identifier] = ctx.ExecutionStack.Pop()
}
