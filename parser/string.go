package parser

func String(content string) Parser[string] {
	return func(input []string) (string, []string, bool) {
		if len(input) == 0 {
			return "", input, false
		}
		
		if input[0] == content {
			return content, input[1:], true
		}
		
		return "", input, false
	}
}