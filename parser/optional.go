package parser

func optional[T any](parser Parser[T]) Parser[T] {
	return func(input []string) (T, []string, bool) {
		result, rest, occured := parser(input)
		
		if occured {
			return result, rest, true
		}
		
		var defaultT T
		return defaultT, input, true
	}
}
