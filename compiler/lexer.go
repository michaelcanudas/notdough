package compiler

type Token struct {
	Value string
	Kind  TokenKind
}

type TokenKind string

const (
	Identifier TokenKind = "_identifier" // underscore means it's a special case and we don't use this value
	Literal              = "_literal"

	OpenParenthesis  = "("
	CloseParenthesis = ")"
	OpenBracket      = "["
	CloseBracket     = "]"
	OpenBrace        = "{"
	CloseBrace       = "}"
	Colon            = ":"
	Semicolon        = ";"
	Comma            = ","
	Dot              = "."
	Quote            = `"`

	Let   = "let"
	Var   = "var"
	If    = "if"
	For   = "for"
	While = "while"

	Equals       = "="
	EqualsEqauls = "=="
	Plus         = "+"
	PlusPlus     = "++"
	Minus        = "-"
	MinusMinus   = "--"
	Star         = "*"
	Slash        = "/"
)

func lex(src string) []Token {
	panic("uh oh (not implemented)")
}
