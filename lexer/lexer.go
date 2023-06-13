package lexer

import (
	"strings"
)

var knownTokens = []string{
	">>=",
	"<=",
	">=",
	"++",
	"==",
	"--",
	"!=",
	">",
	"<",
	"(",
	")",
	"[",
	"]",
	"{",
	"}",
	":",
	";",
	",",
	".",
	`"`,
	"!",
	"=",
	"+",
	"-",
	"*",
	"/",
}

func Lex(src string) []string {
	var (
		currentPos = 0
		currentTok = 0
	)

	var tokens []string

	for currentPos < len(src) {
		if currentPos == currentTok {
			currentPos++
			continue
		}

		tok := src[currentTok:currentPos]

		if tok[0] == '"' {
			if len(tok) > 1 && strings.HasSuffix(tok, `"`) && !strings.HasSuffix(tok, `\"`) {
				tokens = append(tokens, tok)
				currentTok = currentPos
			}

			currentPos++
			continue
		}

		if isWhitespace(tok[0]) {
			currentTok++
			continue
		}

		isTokLetter := isIdentifierChar(tok[0])
		isNextTokLetter := isIdentifierChar(tok[len(tok)-1])

		if isWhitespace(tok[len(tok)-1]) || (isNextTokLetter != isTokLetter) {
			value := src[currentTok : currentPos-1]
			tokens = append(tokens, value)
			currentTok = currentPos - 1
			continue
		}

		// try to match operators, we can only be sure when we have one match
		match, matchCount := "", 0
		for _, kt := range knownTokens {
			if strings.HasPrefix(kt, tok) {
				matchCount++
				match = kt
			}
		}

		if matchCount == 1 && match == tok {
			tokens = append(tokens, match)
			currentTok = currentPos
			continue
		}

		currentPos++
	}

	last := src[currentTok:currentPos]
	last = strings.Trim(last, " ")
	last = strings.Trim(last, "\r")
	last = strings.Trim(last, "\n")
	last = strings.Trim(last, "\t")

	if len(last) > 0{
		tokens = append(tokens, last)
	}

	return tokens
}

func isIdentifierChar(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\r' || char == '\n' || char == '\t'
}
