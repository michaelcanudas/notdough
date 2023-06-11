package interpreter

import "strconv"

type Push struct {
	Value int64
}

func init() {
	RegisterInstruction("push", func(fields []string) Instruction {
		value, _ := strconv.ParseInt(fields[1], 10, 64)
		return Push{
			Value: value,
		}
	})
}

func (p Push) execute(ctx *Context) {
	ctx.ExecutionStack.Push(p.Value)
}
