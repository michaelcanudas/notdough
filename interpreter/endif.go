package interpreter

type Endif struct {
}

func init() {
	RegisterInstruction("endif", func(fields []string) Instruction {
		return Endif{}
	})
}

func (e Endif) execute(ctx *Context) {
	// nop
}
