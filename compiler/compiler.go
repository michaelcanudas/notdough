package compiler

import (
	"strconv"
	"strings"

	"michaelcanudas.dough/interpreter"
)

func Compile(text string) []interpreter.Instruction {
	const bitwiseKeyword string = "bits" // should this be "bitwise" instead ?

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
		case "neg":
			instructions = append(instructions, interpreter.Neg{})
		case "mul":
			instructions = append(instructions, interpreter.Mul{})
		case "div":
			instructions = append(instructions, interpreter.Div{})
		case "rem":
			instructions = append(instructions, interpreter.Rem{})
		case "dup":
			instructions = append(instructions, interpreter.Dup{})
		case "read":
			instructions = append(instructions, interpreter.Read{})
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
		case "cmp":
			instructions = append(instructions, interpreter.Cmp{
				Comparison: fields[1],
			})
		case "and":
			instructions = append(instructions, interpreter.And{
				Bitwise: len(fields) >= 2 && fields[1] == bitwiseKeyword,
			})
		case "or":
			instructions = append(instructions, interpreter.Or{
				Bitwise: len(fields) >= 2 && fields[1] == bitwiseKeyword,
			})
		case "xor":
			instructions = append(instructions, interpreter.Xor{
				Bitwise: len(fields) >= 2 && fields[1] == bitwiseKeyword,
			})
		case "not":
			instructions = append(instructions, interpreter.Not{
				Bitwise: len(fields) >= 2 && fields[1] == bitwiseKeyword,
			})
		case "store":
			instructions = append(instructions, interpreter.Store{
				Identifier: fields[1],
			})
		case "load":
			instructions = append(instructions, interpreter.Load{
				Identifier: fields[1],
			})
		default:
			instructions = append(instructions, interpreter.Call{
				Identifier: term,
			})
		}
	}

	return instructions
}
