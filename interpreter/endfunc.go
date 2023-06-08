package interpreter

type Endfunc struct {
}

func init() {
	RegisterInstruction("endfunc", func(fields []string) Instruction {
		return Endfunc{}
	})
}

func (e Endfunc) execute(ctx *Context) {
	// 'endfunc' has the same behavior as 'Ret'
	ret := Ret{}
	ret.execute(ctx)
}
