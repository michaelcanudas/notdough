package interpreter

import (
	"fmt"
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/lexer"
	"michaelcanudas.dough/parser"
)


type InstructionProvider func(node ast.Node)Instruction

var registeredInstructionProviders = make(map[string]InstructionProvider)

func IlCompileString(src string) []Instruction {
	return IlCompile(parser.IlParse(lexer.Lex(src)))
}

func IlCompile(nodes []ast.IlInstructionNode) []Instruction {
	var instructions = make([]Instruction, len(nodes))
	for i, n := range nodes {
		instructions[i] = IlCompileNode(n)
	}
	return instructions
}

func IlCompileNode(node ast.IlInstructionNode) Instruction {
	provider, ok := registeredInstructionProviders[node.OpCode.Value]

	if !ok {
		provider, ok = registeredInstructionProviders[""]
	}

	if !ok {
		panic(fmt.Sprintf("unknown instruction '%v'", node.OpCode.Value))
	}

	return provider(node.Argument)
}

func RegisterInstruction(name string, provider InstructionProvider) {
	registeredInstructionProviders[name] = provider
}