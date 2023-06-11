package interpreter

type Instruction interface {
	execute(*Context)
}

type NamedInstruction interface {
	getName() string
}

type Context struct {
	// when we add variables:
	// push a map on function call
	// pop it on function return

	LocalStack     Stack[map[string]int64] // stack frames for each currently executing function
	ExecutionStack Stack[int64]            // current stack of operands for direct use by instructions
	ReturnStack    Stack[int64]            // the stack of return locations, used by call/ret
	FuncMap        map[string]int64        // the location of each instruction
	Instructions   []Instruction           // the current program's instructions
	Pointer        int64                   // The current instruction pointer
}
