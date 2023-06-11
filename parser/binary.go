package parser

type BinaryExpressionNode struct {
	Op string
	Lhs ExpressionNode
	Rhs ExpressionNode
}

func Add(lhs ExpressionNode) Parser[BinaryExpressionNode] {
	return func(input []string) (BinaryExpressionNode, []string, bool) {
		op, rest, ok := String("+")(input)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		term, rest, ok := Term()(rest)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		return BinaryExpressionNode {
			Op: op,
			Lhs: lhs,
			Rhs: term,
		}, rest, ok
	}
}

func Sub(lhs ExpressionNode) Parser[BinaryExpressionNode] {
	return func(input []string) (BinaryExpressionNode, []string, bool) {
		op, rest, ok := String("-")(input)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		term, rest, ok := Term()(rest)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		return BinaryExpressionNode {
			Op: op,
			Lhs: lhs,
			Rhs: term,
		}, rest, ok
	}
}

func Mul(lhs ExpressionNode) Parser[BinaryExpressionNode] {
	return func(input []string) (BinaryExpressionNode, []string, bool) {
		op, rest, ok := String("*")(input)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		factor, rest, ok := Factor()(rest)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		return BinaryExpressionNode {
			Op: op,
			Lhs: lhs,
			Rhs: factor,
		}, rest, ok
	}
}

func Div(lhs ExpressionNode) Parser[BinaryExpressionNode] {
	return func(input []string) (BinaryExpressionNode, []string, bool) {
		op, rest, ok := String("/")(input)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		factor, rest, ok := Factor()(rest)
		if !ok {
			return BinaryExpressionNode{}, input, ok
		}

		return BinaryExpressionNode {
			Op: op,
			Lhs: lhs,
			Rhs: factor,
		}, rest, ok
	}
}
