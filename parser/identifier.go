package parser

import "regexp"

func Identifier() Parser[string] {
	return func(input []string) (string, []string, bool) {
		match, _ := regexp.MatchString("^[_a-zA-Z][_a-zA-Z0-9]*$", input[0])
		
		result := ""
		rest := input
		if match {
			result = input[0]
			rest = input[1:]
		}
		
		return result, rest, match
	}
}