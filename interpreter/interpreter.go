package interpreter

func Interpret(instructions []Instruction) {
	context := Context{
		LocalStack:     Stack[map[string]int64]{make(map[string]int64)},
		ExecutionStack: Stack[int64]{},
		ReturnStack:    Stack[int64]{},
		FuncMap:        map[string]int64{},
		Instructions:   instructions,
		Pointer:        0,
	}

	for i := 0; i < len(context.Instructions); i++ {
		instruction, ok := context.Instructions[i].(NamedInstruction)
		if ok {
			context.FuncMap[instruction.getName()] = int64(i)
		}
	}

	for context.Pointer < int64(len(context.Instructions)) {
		context.Instructions[context.Pointer].execute(&context)
		context.Pointer++
	}
}

type InstructionProviderMap map[string]func(fields []string) Instruction

var RegisteredInstructionProviders InstructionProviderMap

func RegisterInstruction(name string, createInstruction func(fields []string) Instruction) {
	if RegisteredInstructionProviders == nil {
		RegisteredInstructionProviders = make(InstructionProviderMap)
	}

	RegisteredInstructionProviders[name] = createInstruction
}
