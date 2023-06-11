package interpreter

import "fmt"

type Print struct {
}

func init() {
	RegisterInstruction("print", func(fields []string) Instruction {
		return Print{}
	})
}

func (p Print) execute(ctx *Context) {
	fmt.Println(ctx.ExecutionStack.Pop())
}
