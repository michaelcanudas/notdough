package commands

import (
	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
	"michaelcanudas.dough/lexer"
	"michaelcanudas.dough/parser"
)

func Build(args []string) ([]interpreter.Instruction, error) {
	if len(args) == 0 {
		panic("no input file found")
	}

	input, err := GetText(args[0])
	if err != nil {
		return nil, err
	}

	_, err = GetSettings(args[1:])
	if err != nil {
		return nil, err
	}

	tokens := lexer.Lex(input)
	nodes := parser.Parse(tokens)
	instructions := compiler.Compile(nodes)

	return instructions, nil
}
