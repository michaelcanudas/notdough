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

		line := lines[i]

		// we probably want to support multiple quoted strings per line, but this is easier :P
		// treat the characters between the first and last quote in the string as one field
		start := strings.Index(line, "\"")
		end := strings.LastIndex(line, "\"")

		var fields []string
		if start != -1 && end != -1 {
			beforeQuotes := strings.Fields(line[:start])
			afterQuotes := strings.Fields(line[end+1:])
			quotes := line[start : end+1]

			fields = append(append(beforeQuotes, quotes), afterQuotes[:]...)
		} else {
			fields = strings.Fields(line)
		}

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
