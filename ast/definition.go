package ast

type DefinitionNode struct {
	Type TypeNode

	Keyword    KeywordNode
	Identifier IdentifierNode
	Assignment SymbolNode
	Value      Node
}
