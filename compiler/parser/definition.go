package parser

type Definition struct {
	Keyword Token
	Identifier Identifier
	Assignment Token
	Expression Expression
}

func parseDefinition(source *[]string) Definition {
	var definition Definition

	definition.Keyword = parseToken(source, "let")
	definition.Identifier = parseIdentifier(source)
	definition.Assignment = parseToken(source, "=")
	definition.Expression = parseExpression(source)

	return definition
}

// let [id] = [ex]