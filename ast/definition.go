package ast

type DefinitionNode struct {
	Keyword KeywordNode
	Identifier IdentifierNode
	Type TypeNode
	Assignment SymbolNode
	Value      Node
}
