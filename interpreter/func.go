package interpreter

type Func struct {
	Identifier string
}

func init() {
	RegisterInstruction("func", func(fields []string) Instruction {
		return Func{
			Identifier: fields[1],
		}
	})
}

func (f Func) execute(ctx *Context) {
	// functions, when executed, just skip to their corresponding 'endfunc'
	// instruction (so that they are not executed until they are called)
	ctx.Pointer = FindEndfunc(ctx)
}
