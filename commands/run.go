package commands

import (
	"michaelcanudas.dough/compiler"
	"michaelcanudas.dough/interpreter"
	"michaelcanudas.dough/lexer"
	"michaelcanudas.dough/parser"
)

func Run(args []string) error {
	if len(args) == 0 {
		panic("no input file found")
	}

	input, err := GetText(args[0])
	if err != nil {
		return err
	}

	_, err = GetSettings(args[1:])
	if err != nil {
		return err
	}

	tokens := lexer.Lex(input)
	nodes := parser.Parse(tokens)
	instructions := compiler.Compile(nodes)

	/*for _, instruction := range instructions {
		fmt.Print(reflect.TypeOf(instruction).Name())
		fmt.Println(instruction)
	}*/

	interpreter.Interpret(instructions)

	return nil
}
