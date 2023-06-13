package parser

func many[T any](parser Parser[T]) Parser[[]T] {
	return func(input []string) ([]T, []string, bool) {
		next := input
		
		var results []T
		for {
			result, rest, ok := parser(next)
			if !ok {
				break
			}
			
			next = rest
			results = append(results, result)
		}
		
		return results, next, true
	}
}
