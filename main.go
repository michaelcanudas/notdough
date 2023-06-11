package main

import (
	"fmt"
	"michaelcanudas.dough/tests"
	"os"
	"path"

	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
)

func main() {
	if len(os.Args) >= 3 && os.Args[1] == "-t" {
		if os.Args[2] == "--all" {
			tests.InvokeAllTests()
		} else {
			tests.InvokeTest(os.Args[2])
		}
		return
	}

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
