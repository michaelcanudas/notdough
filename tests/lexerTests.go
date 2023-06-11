package tests

import (
	"fmt"
	"michaelcanudas.dough/compiler"
)

func init() {
	RegisterTest("lexer.main", func() {
		txt := `import system

let main = void() {
loop(10, void() {
system.print("hello!")
})
}

let loop = void(x: int, fn: void()) {
if x <= 0 {
return
}

fn()
loop(x - 1, fn)
}`
		lexed := compiler.Lex(txt)

		for _, l := range lexed {
			fmt.Println("_" + l + "_")
		}
	})
}
