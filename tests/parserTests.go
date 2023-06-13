package tests

import (
	"fmt"
	"strings"

	"michaelcanudas.dough/lexer"
	"michaelcanudas.dough/parser"
)

func init() {
	RegisterTest("parser.main", func() {
		input := `
let main: int = {
	let x = 20 + 30 * 40
	return x
}
`

		tokens := lexer.Lex(strings.TrimSpace(input))
		node := parser.Parse(tokens)

		if node == nil {
			panic("failed to parse")
		}

		fmt.Println(node)
	})
}
