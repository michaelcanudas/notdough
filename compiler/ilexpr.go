package compiler

import (
	"michaelcanudas.dough/ast"
	"michaelcanudas.dough/interpreter"
)

func ilexpr(expr ast.IlExpression)[]interpreter.Instruction {
	return interpreter.IlCompile(expr.Instructions)
}