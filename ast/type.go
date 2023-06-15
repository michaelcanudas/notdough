package ast

type KeywordTypeNode struct {
	Type KeywordNode
}

type FunctionTypeNode struct {
	Type     TypeNode
	ArgTypes []TypeNode
}
