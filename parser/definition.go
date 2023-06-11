package parser

type DefinitionNode struct {
	Identifier string
}

func Definition() Parser[DefinitionNode] {
	return func(input []string) (DefinitionNode, []string, bool) {
		node := DefinitionNode{}
		
		_, rest, ok := String("let")(input)
		if !ok {
			return node, input, ok
		}
		
		identifier, rest, ok := Identifier()(rest)
		if !ok {
			return node, input, ok
		}
		
		return DefinitionNode{
			Identifier: identifier,
		}, rest, ok
	}
}
