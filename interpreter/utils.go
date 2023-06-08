package interpreter

type Stack[T any] []T

// Push pushes an element to the top of a stack.
func (stack *Stack[T]) Push(element T) {
	*stack = append(*stack, element)
}

// Pop removes and returns the topmost element of a stack.
func (stack *Stack[T]) Pop() T {
	var index = len(*stack) - 1  // index of topmost element
	var result = (*stack)[index] // get result value
	*stack = (*stack)[:index]    // shrink stack
	return result
}

// Peek returns the topmost element of a stack without removing it.
func (stack *Stack[T]) Peek() T {
	var index = len(*stack) - 1
	var result = (*stack)[index]
	return result
}

func FindEndif(ctx *Context) int64 {
	loc := ctx.Pointer
	depth := 1

	for depth > 0 {
		loc++

		switch ctx.Instructions[loc].(type) {
		case If:
			depth++
		case Endif:
			depth--
		}
	}

	return loc
}

func FindEndfunc(ctx *Context) int64 {
	loc := ctx.Pointer
	depth := 1

	for depth > 0 {
		loc++

		switch ctx.Instructions[loc].(type) {
		case Func:
			depth++
		case Endfunc:
			depth--
		}
	}

	return loc
}

func FindLoop(ctx *Context) int64 {
	loc := ctx.Pointer
	depth := 1

	for depth > 0 {
		loc--

		switch ctx.Instructions[loc].(type) {
		case Loop:
			depth--
		case Endloop:
			depth++
		}
	}

	return loc
}
