package parser

func Optional[T any](parser Parser[T]) Parser[T] {
	return func(input []string) (T, []string, bool) {
		result, rest, ok := parser(input)
		
		if ok {
			return result, rest, ok
		}
		
		var defaultT T
		return defaultT, input, true
	}
}
