package tests

import (
	"strings"

	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
	"michaelcanudas.dough/lexer"
	"michaelcanudas.dough/parser"
)

func init() {
	RegisterTest("compiler.main", func() {
		input := `
{
print 1 - 2 + 3
}
`

		tokens := lexer.Lex(strings.TrimSpace(input))
		node := parser.Parse(tokens)

		if node == nil {
			panic("parse error")
		}

		instructions := compiler.Compile(node)

		if instructions == nil {
			panic("compiler error")
		}

		interpreter.Interpret(instructions)
	})
}
