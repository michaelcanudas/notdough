package interpreter

type Instruction interface {
	execute(*Context)
}

type Context struct {
	ExecutionStack Stack[int64]
	ReturnStack    Stack[int64]
	FuncMap        map[string]int64
	Instructions   []Instruction
	Pointer        int64
}
