package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func Compile(node ast.Node) []interpreter.Instruction {
	switch node.(type) {
	case ast.BinaryNode:
		return binary(node.(ast.BinaryNode))
	case ast.BlockNode:
		return block(node.(ast.BlockNode))
	case ast.CallNode:
		return call(node.(ast.CallNode))
	case ast.DefinitionNode:
		return definition(node.(ast.DefinitionNode))
	case ast.IdentifierNode:
		return identifier(node.(ast.IdentifierNode))
	case ast.NumberNode:
		return number(node.(ast.NumberNode))
	case ast.PrintNode:
		return print(node.(ast.PrintNode))
	case ast.ReturnNode:
		return ret(node.(ast.ReturnNode))
	case ast.UnaryNode:
		return unary(node.(ast.UnaryNode))
	case ast.IlExpression:
		return ilexpr(node.(ast.IlExpression))
	default:
		return nil
	}
}
