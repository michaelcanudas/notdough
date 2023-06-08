package main

import (
	"os"

	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
)

func main() {
	input, _ := os.ReadFile("./input.txt")

	interpreter.Interpret(compiler.Compile(string(input)))
}
