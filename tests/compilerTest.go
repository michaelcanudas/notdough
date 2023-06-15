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
	let x = {
		5 * 2
	}
	let y = 25
	print (x + y) + 5

	print il {
		push 70
		sub
	}

	print 1 + 2 + 3
}
`

		tokens := lexer.Lex(strings.TrimSpace(input))
		node := parser.Parse(tokens)

		if node == nil {
			panic("parse failed")
		}

		instructions := compiler.Compile(node)

		if instructions == nil {
			panic("compilation failed")
		}

		interpreter.Interpret(instructions)
	})
}
