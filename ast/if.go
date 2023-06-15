package ast

type If struct {
	Condition Node
	Body      Node
	ElseBlock Node
}
