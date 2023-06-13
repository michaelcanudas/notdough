package interpreter

import (
	"michaelcanudas.dough/ast"
	"strings"
)

func init() {
	RegisterInstruction("rsrc", func(arg ast.Node) Instruction {
		return Rsrc{
			Name: arg.(ast.IdentifierNode).Value,
		}
	})
}

func unescape(s string) string {
	r := s[1 : len(s)-1]
	r = strings.Replace(r, `\n`, "\n", -1)
	r = strings.Replace(r, `\"`, "\"", -1)
	return r
}

type Rsrc struct {
	Name    string
}

func (r Rsrc) execute(ctx *Context) {
	// nop
}

func (r Rsrc) getName() string {
	return r.Name
}