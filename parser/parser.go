package parser

type Parser[T any] func([]string) (T, []string, bool)
