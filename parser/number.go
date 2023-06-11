package parser

type NumberExpressionNode struct {
	Value int64
}

func Number() Parser[NumberExpressionNode] {
	return func(input []string) (NumberExpressionNode, []string, bool) {
		if len(input) == 0 {
			return NumberExpressionNode{}, input, false
		}

		if x, err := strconv.ParseInt(input[0], 10, 64); err == nil {
			return NumberExpressionNode {
				Value: x,
			}, input[1:], true
		}

		return NumberExpressionNode{}, input, false
	}
}
