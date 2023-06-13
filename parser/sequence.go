package parser

func sequence[T any](parsers ...Parser[T]) Parser[[]T] {
	return func(input []string) ([]T, []string, bool) {
		next := input
		
		var results []T
		for i := 0; i < len(parsers); i++ {
			result, rest, ok := parsers[i](next)
			if !ok {
				return nil, input, ok
			}

			next = rest
			results = append(results, result)
		}

		return results, next, true
	}
}

/*
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
*/
