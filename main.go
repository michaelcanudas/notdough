package main

import (
	"fmt"
	"os"
	"path"

	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
)

func main() {
	var file string
	if len(os.Args) >= 2 {
		file = os.Args[1]
	} else {
		file = "input.txt"
	}

	input, err := os.ReadFile(file)

	if err != nil {
		input, err = os.ReadFile(path.Join("ilcode/", file))

		if err != nil {
			panic(fmt.Sprintf("unable to read code file '%s'!!", file))
		}
	}

	interpreter.Interpret(compiler.Compile(string(input), interpreter.RegisteredInstructionProviders))
}
