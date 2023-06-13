package parser

import (
	"fmt"
	"michaelcanudas.dough/ast"
)


func instruction() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		opcode, rest, ok := identifier()(input)
		
		if !ok {
			return nil, input, false
		}
		
		opcodeIdentifier, ok := opcode.(ast.IdentifierNode)
		
		if !ok {
			return nil, input, false
		}
		
		arg, rest, ok := optional(either(
			number(),
			literal(),
			symbol("<"),
			symbol(">"),
			symbol("!="),
			symbol("=="),
			symbol("<="),
			symbol(">="),
			ilvar(),
			))(rest)
		
		if !ok {
			return nil, input, false
		}
		
		node := ast.IlInstructionNode{
			OpCode: opcodeIdentifier,
			Argument: arg,
		}
		
		return node, rest, true
	}
}

func ilvar() Parser[ast.Node] {
	return func(input []string) (ast.Node, []string, bool) {
		_, rest, ok := symbol("$")(input)
		
		if !ok {
			return nil, input, false
		}
		
		name, rest, ok := identifier()(rest)
		
		nameIdentifier, ok := name.(ast.IdentifierNode)
		if !ok {
			return nil, input, false
		}
		
		result := ast.IlVariableNode{
			VarName: nameIdentifier,
		}
		return result, rest, true
	}
}

func ilExpression() Parser[ast.Node] {
	return either(func(input []string) (ast.Node, []string, bool) {
		_, rest, ok := keyword("il")(input)

		if !ok {
			return nil, input, false
		}

		_, rest, ok = symbol("{")(rest)

		if !ok {
			return nil, input, false
		}

		fmt.Println(rest)
		insts, rest, _ := many(instruction())(rest)

		instructions := []ast.IlInstructionNode{}
		
		if insts == nil {
			instructions = nil
		} else {
			for	_, i := range insts {
				instructions = append(instructions, i.(ast.IlInstructionNode))
			}
		}
		
		_, rest, ok = symbol("}")(rest)

		if !ok {
			return nil, input, false
		}

		result := ast.IlExpression{
			Instructions: instructions,
			}

			return result, rest, true
	}, print());
}

// feel free to move this
func cast[To ast.Node](p Parser[ast.Node]) Parser[To] {
	return func(input []string) (To, []string, bool) {
		f, rest, ok := p(input)
		return f.(To), rest, ok
	}
}