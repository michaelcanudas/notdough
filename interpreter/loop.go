package interpreter

type Loop struct {
}

func (l Loop) execute(ctx *Context) {
	// nop
	// loops only serve as a header for other instructions (such as cont) to jump to
}
