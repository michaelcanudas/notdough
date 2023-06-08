package compiler

import (
	"strings"

	"michaelcanudas.dough/interpreter"
)

func Compile(text string, instructionProviders interpreter.InstructionProviderMap) []interpreter.Instruction {

	lines := strings.Split(text, "\n")

	var instructions []interpreter.Instruction
	for i := 0; i < len(lines); i++ {
		trimmed := strings.Trim(lines[i], " ")
		if len(trimmed) == 0 || trimmed[0] == '#' {
			continue
		}

		fields := strings.Fields(lines[i])

		if len(fields) == 0 {
			continue
		}

		var fac = instructionProviders[fields[0]]
		if fac == nil {
			fac = instructionProviders[""]
		}
		if fac == nil {
			panic("unknown instruction and unknown handler (call instruction) not registered!")
		}

		var inst = fac(fields)
		instructions = append(instructions, inst)
	}

	return instructions
}
