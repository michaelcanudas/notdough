package interpreter

type Cont struct {
}

func init() {
	RegisterInstruction("cont", func(fields []string) Instruction {
		return Cont{}
	})
}
func (c Cont) execute(ctx *Context) {
	// continue instruction skips to the next iteration of the current loop

	// find the loop instruction
	loop := FindLoop(ctx)
	// and go to it
	ctx.Pointer = loop
}
