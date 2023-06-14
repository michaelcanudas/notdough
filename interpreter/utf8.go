package interpreter

import (
	"michaelcanudas.dough/ast"
	"strings"
)

type Text struct {
	Value string
}

func init() {
	RegisterInstruction("text", func(node ast.Node) Instruction {
		return Text{
			Value: unescape(node.(ast.LiteralNode).Value),
		}
	})
}

func (t Text) execute(ctx *Context) {
	// nop
}

func unescape(s string) string {
	r := strings.Replace(s, `\n`, "\n", -1)
	r = strings.Replace(r, `\"`, "\"", -1)
	return r
}
