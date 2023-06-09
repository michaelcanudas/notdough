package parser

func Parse(source []string) Program {
	parseSource := source
	
	return parseProgram(&parseSource)
}