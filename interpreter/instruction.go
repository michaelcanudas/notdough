package interpreter

type Instruction interface {
	execute(*Context)
}

type Context struct {
	Stack []int64
	ReturnStack []int64
	FuncMap map[string]int64
	Instructions []Instruction
	Pointer int64
}