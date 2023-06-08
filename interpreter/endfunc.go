package interpreter

type Endfunc struct {
}

func (e Endfunc) execute(ctx *Context) {
	ret := Ret{}
	ret.execute(ctx)
}
