package Lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return rune(l.input[l.pos])
}

func (l *Lexer) advance() {
	l.pos++
}

func (l *Lexer) NextToken() Token {
	for l.peek() != 0 {
		ch := l.peek()

		if unicode.IsSpace(ch) {
			l.advance()
			continue
		}

		if unicode.IsDigit(ch) {
			start := l.pos
			for unicode.IsDigit(l.peek()) {
				l.advance()
			}
			return Token{Type: NUMBER, Value: l.input[start:l.pos]}
		}

		switch ch {
		case '+':
			l.advance()
			return Token{Type: PLUS, Value: "+"}
		case '-':
			l.advance()
			return Token{Type: MINUS, Value: "-"}
		case '*':
			l.advance()
			return Token{Type: STAR, Value: "*"}
		case '/':
			l.advance()
			return Token{Type: SLASH, Value: "/"}
		default:
			panic(fmt.Sprintf("Unknow character: %q", ch))
		}
	}
	return Token{Type: EOF, Value: ""}
}

func (l *Lexer) Lex() []Token {
	var tokens []Token
	for {
		tok := l.NextToken()
		if tok.Type == EOF {
			break
		}
		tokens = append(tokens, tok)
	}
	return tokens
}
