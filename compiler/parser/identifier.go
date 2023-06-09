package parser

type Identifier struct {
	Value string
}

func parseIdentifier(source *[]string) Identifier {
	var identifier Identifier

	identifier.Value = (*source)[0]
	*source = (*source)[1:]

	return identifier
}