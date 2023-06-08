package compiler

import (
	"strconv"
	"strings"

	"michaelcanudas.dough/interpreter"
)

func Compile(text string) []interpreter.Instruction {
	lines := strings.Split(text, "\n")

	var instructions []interpreter.Instruction
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		if strings.Trim(lines[i], " ")[0] == '#' {
			continue
		}

		fields := strings.Fields(lines[i])

		if len(fields) == 0 {
			continue
		}

		term := fields[0]

		switch term {
		case "push":
			value, _ := strconv.ParseInt(fields[1], 10, 64)
			instructions = append(instructions, interpreter.Push{
				Value: value,
			})
		case "pop":
			value, _ := strconv.ParseInt(fields[1], 10, 64)
			instructions = append(instructions, interpreter.Pop{
				Count: value,
			})
		case "print":
			instructions = append(instructions, interpreter.Print{})
		case "add":
			instructions = append(instructions, interpreter.Add{})
		case "sub":
			instructions = append(instructions, interpreter.Sub{})
		case "mul":
			instructions = append(instructions, interpreter.Mul{})
		case "if":
			instructions = append(instructions, interpreter.If{})
		case "endif":
			instructions = append(instructions, interpreter.Endif{})
		case "loop":
			instructions = append(instructions, interpreter.Loop{})
		case "cont":
			instructions = append(instructions, interpreter.Cont{})
		case "endloop":
			instructions = append(instructions, interpreter.Endloop{})
		case "func":
			instructions = append(instructions, interpreter.Func{
				Identifier: fields[1],
			})
		case "endfunc":
			instructions = append(instructions, interpreter.Endfunc{})
		case "ret":
			instructions = append(instructions, interpreter.Ret{})
		default:
			instructions = append(instructions, interpreter.Call{
				Identifier: term,
			})
		}
	}

	return instructions
}
