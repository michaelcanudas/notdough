package parser

import "strconv"

func Expression() Parser[int] {
	return func(input []string) (int, []string, bool) {
		return Either(func(input []string) (int, []string, bool) {
			term, rest, ok := Term()(input)
			if !ok {
				return 0, input, ok
			}
			
			return Either(Add(term), Sub(term))(rest)
		}, Term())(input)
	}
}

func Add(lhs int) Parser[int] {
	return func(input []string) (int, []string, bool) {
		_, rest, ok := String("+")(input)
		if !ok {
			return 0, input, ok
		}
		
		term, rest, ok := Term()(rest)
		if !ok {
			return 0, input, ok
		}
		
		return lhs + term, rest, ok
	}
}

func Sub(lhs int) Parser[int] {
	return func(input []string) (int, []string, bool) {
		_, rest, ok := String("-")(input)
		if !ok {
			return 0, input, ok
		}
		
		term, rest, ok := Term()(rest)
		if !ok {
			return 0, input, ok
		}

		return lhs - term, rest, ok
	}
}

func Term() Parser[int] {
	return func(input []string) (int, []string, bool) {
		return Either(func(input []string) (int, []string, bool) {
			factor, rest, ok := Factor()(input)
			if !ok {
				return 0, input, false
			}

			return Either(Mul(factor), Div(factor))(rest)
		}, Factor())(input)
	}
}

func Mul(lhs int) Parser[int] {
	return func(input []string) (int, []string, bool) {
		_, rest, ok := String("*")(input)
		if !ok {
			return 0, input, ok
		}
		
		factor, rest, ok := Factor()(rest)
		if !ok {
			return 0, input, ok
		}

		return lhs * factor, rest, ok
	}
}

func Div(lhs int) Parser[int] {
	return func(input []string) (int, []string, bool) {
		_, rest, ok := String("/")(input)
		if !ok {
			return 0, input, ok
		}
		
		factor, rest, ok := Factor()(rest)
		if !ok {
			return 0, input, ok
		}

		return lhs / factor, rest, ok
	}
}

func Factor() Parser[int] {
	return func(input []string) (int, []string, bool) {
		return Either(Digit(), func(input []string) (int, []string, bool) {
			_, rest, ok := String("(")(input)
			if !ok {
				return 0, input, ok
			}
			
			expression, rest, ok := Expression()(rest)
			if !ok {
				return 0, input, ok
			}
			
			_, rest, ok = String(")")(rest)
			if !ok {
				return 0, input, ok
			}
			
			return expression, rest, ok
		})(input)
	}
}

func Digit() Parser[int] {
	return func(input []string) (int, []string, bool) {
		if len(input) == 0 {
			return 0, input, false
		}
		
		if x, err := strconv.Atoi(input[0]); err == nil {
			return x, input[1:], true
		}
		
		return 0, input, false
	}
}