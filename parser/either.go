package parser

func either[T any](parsers ...Parser[T]) Parser[T] {
	return func(input []string) (T, []string, bool) {
		for i := 0; i < len(parsers); i++ {
			result, rest, ok := parsers[i](input)

			if ok {
				return result, rest, ok
			}
		}

		var result T
		return result, input, false
	}
}
