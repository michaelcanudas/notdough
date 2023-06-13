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

	1 + 2
	il {
		push 4
		mul
		print
	}
}
`

		tokens := lexer.Lex(strings.TrimSpace(input))
		node := parser.Parse(tokens)
		instructions := compiler.Compile(node)
		
		interpreter.Interpret(instructions)
	})
}
