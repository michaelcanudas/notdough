package ast

type IlInstructionNode struct {
	OpCode   IdentifierNode
	Argument Node // may be nil
}

type IlVariableNode struct {
	VarName IdentifierNode
}

type IlExpression struct {
	Instructions []IlInstructionNode
}
