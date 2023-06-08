package interpreter

import "strconv"

type Pop struct {
	Count int64
}

func init() {
	RegisterInstruction("pop", func(fields []string) Instruction {
		count, _ := strconv.ParseInt(fields[1], 10, 64)
		return Pop{
			Count: count,
		}
	})
}

func (p Pop) execute(ctx *Context) {
	ctx.ExecutionStack.Pop()
}
