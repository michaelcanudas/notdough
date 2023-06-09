package parser

import "errors"

type Token struct {
	Value string
}

func parseToken(source *[]string, value string) Token {
	var token Token

	if (*source)[0] != value {
		panic(errors.New("invalid token"))
	}

	token.Value = value
	*source = (*source)[1:]

	return token
}