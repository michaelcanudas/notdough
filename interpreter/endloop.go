package interpreter

type Endloop struct {
}

func init() {
	RegisterInstruction("endloop", func(fields []string) Instruction {
		return Endloop{}
	})
}

func (e Endloop) execute(ctx *Context) {
	// nop
}
