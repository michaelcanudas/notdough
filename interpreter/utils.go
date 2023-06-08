package interpreter

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
