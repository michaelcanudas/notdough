package interpreter

type Instruction interface {
	execute(*Context)
}

type Context struct {
	// when we add variables:
	// push a map on function call
	// pop it on function return
	//
	// LocalsStack    Stack[map[string]int64]

	ExecutionStack Stack[int64]
	ReturnStack    Stack[int64]
	FuncMap        map[string]int64
	Instructions   []Instruction
	Pointer        int64
}
