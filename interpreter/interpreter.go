package interpreter

func Interpret(instructions []Instruction) {
	context := Context{
		ReturnStack:    Stack[int64]{},
		ExecutionStack: Stack[int64]{},
		FuncMap:        map[string]int64{},
		Instructions:   instructions,
		Pointer:        0,
	}

	for i := 0; i < len(context.Instructions); i++ {
		var f, ok = context.Instructions[i].(Func)
		if ok {
			context.FuncMap[f.Identifier] = int64(i)
		}
	}

	for context.Pointer < int64(len(context.Instructions)) {
		context.Instructions[context.Pointer].execute(&context)
		context.Pointer++
	}
}
