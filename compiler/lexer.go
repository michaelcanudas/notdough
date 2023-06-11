package compiler

import (
	"fmt"
	"michaelcanudas.dough/tests"
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

	tokens = append(tokens, strings.Trim(src[currentTok:currentPos], " "))

	return tokens
}

func init() {
	tests.RegisterTest("lexer", func() {
		txt := `import system

		let main = void() {
		loop(10, void() {
		system.print("hello!")
		})
		}

		let loop = void(x: int, fn: void()) {
		if x <= 0 {
		return
		}

		fn()
		loop(x - 1, fn)
		}`
		lexed := Lex(txt)

		for _, l := range lexed {
			fmt.Println("_" + l + "_")
		}

	})
}

func isIdentifierChar(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\r' || char == '\n' || char == '\t'
}
