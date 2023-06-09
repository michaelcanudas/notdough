package parser

type Program struct {
	Definitions []Definition
}

func parseProgram(source *[]string) Program {
	var program Program

	for len(*source) != 0 {
		program.Definitions = append(program.Definitions, parseDefinition(source))
	}

	return program
}