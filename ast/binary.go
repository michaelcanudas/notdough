package ast

type BinaryNode struct {
	Op  SymbolNode
	Lhs Node
	Rhs Node
}
