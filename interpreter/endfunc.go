package interpreter

type Endfunc struct {
}

func (e Endfunc) execute(ctx *Context) {
	// 'endfunc' has the same behavior as 'Ret'
	ret := Ret{}
	ret.execute(ctx)
}
