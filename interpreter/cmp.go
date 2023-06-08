package interpreter

import "fmt"

type Cmp struct {
	Comparison string
}

func (c Cmp) execute(ctx *Context) {
	// pop operands off execution stack
	var a = ctx.ExecutionStack.Pop()
	var b = ctx.ExecutionStack.Pop()

	// actually do the comparison
	var result bool
	switch c.Comparison {
	case "<":
		result = a < b
	case ">":
		result = a > b
	case ">=":
		result = a >= b
	case "<=":
		result = a <= b
	case "==":
		result = a == b
	case "!=":
		result = a != b
	default:
		panic(fmt.Sprintf("unknown comparison %s!!!!", c.Comparison))
	}

	// there has got to be a better way to do this :/
	var value int64
	if result {
		value = 1
	} else {
		value = 0
	}

	ctx.ExecutionStack.Push(value)
}
