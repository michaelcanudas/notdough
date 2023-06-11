package parser

type ExpressionNode interface {
}

func Expression() Parser[ExpressionNode] {
	return func(input []string) (ExpressionNode, []string, bool) {
		return Either(func(input []string) (ExpressionNode, []string, bool) {
			term, rest, ok := Term()(input)
			if !ok {
				return nil, input, ok
			}

			return Either(Add(term), Sub(term))(rest)
		}, Term())(input)
	}
}

func Term() Parser[ExpressionNode] {
	return func(input []string) (ExpressionNode, []string, bool) {
		return Either(func(input []string) (ExpressionNode, []string, bool) {
			factor, rest, ok := Factor()(input)
			if !ok {
				return nil, input, ok
			}

			return Either(Mul(factor), Div(factor))(rest)
		}, Factor())(input)
	}
}

func Factor() Parser[ExpressionNode] {
	return func(input []string) (ExpressionNode, []string, bool) {
		panic("NOT IMPLEMENTED")
		// return Either[ExpressionNode](Number(), func(input []string) (ExpressionNode, []string, bool) {
		// 		_, rest, ok := String("(")(input)
		// 		if !ok {
		// 			return nil, input, ok
		// 		}
		//
		// 		expression, rest, ok := Expression()(rest)
		// 		if !ok {
		// 			return nil, input, ok
		// 		}
		//
		// 		_, rest, ok = String(")")(rest)
		// 		if !ok {
		// 			return nil, input, ok
		// 		}
		//
		// 		return expression, rest, ok
		// })(input)
	}
}
