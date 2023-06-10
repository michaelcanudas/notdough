package interpreter

import "fmt"

func init() {
	RegisterInstruction("printstr", func(fields []string) Instruction {
		return Printstr{}
	})
}

type Printstr struct {
}

func (p Printstr) execute(ctx *Context) {
	resourceLocation := ctx.ExecutionStack.Pop()
	instruction, ok := ctx.Instructions[resourceLocation].(Rsrc)

	if !ok {
		panic("Printstr did not receive a resource!")
	}

	content, ok := instruction.Content.(StringResourceContent)

	if !ok {
		panic("Not a string resource!")
	}

	fmt.Println(content.Value)
}
