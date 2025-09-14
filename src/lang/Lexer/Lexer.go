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

/*
	Making Tokens
*/

func (l *Lexer) NextToken() Token {
	for l.peek() != 0 {
		ch := l.peek()

		if unicode.IsSpace(ch) {
			l.advance()
			continue
		}

		if unicode.IsDigit(ch) {
			return l.MakeNumber()
		}
		if unicode.IsLetter(ch) {
			return l.MakeText()
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
		case '=':
			l.advance()
			return Token{Type: EQUALS, Value: "="}
		case '(':
			l.advance()
			return Token{Type: LPAREN, Value: "("}
		case ')':
			l.advance()
			return Token{Type: RPAREN, Value: ")"}
		default:
			panic(fmt.Sprintf("Unknow character: %q", ch))
		}
	}
	return Token{Type: EOF, Value: ""}
}

/*
	Number Creation
*/

func (l *Lexer) MakeNumber() Token {
	start := l.pos
	hasDot := false
	for {
		if unicode.IsDigit(l.peek()) {
			l.advance()
		} else if l.peek() == ',' && hasDot == false {
			hasDot = true
			l.advance()
		} else if l.peek() == ',' && hasDot == true {
			panic("Dot not allowed in number")
		} else {
			break
		}
	}
	if hasDot {
		return Token{Type: FLOAT, Value: l.input[start:l.pos]}
	}
	return Token{Type: INT, Value: l.input[start:l.pos]}
}

/*
	Creating Text Tokens
*/

func (l *Lexer) MakeText() Token {
	start := l.pos
	for unicode.IsLetter(l.peek()) {
		l.advance()
	}
	text := l.input[start:l.pos]
	switch text {
	case "let":
		return Token{Type: LET, Value: "let"}
	case "println":
		return Token{Type: PRINTLN, Value: "println"}
	default:
		return Token{Type: IDENTIFIER, Value: text}
	}
}

/*
Main Lexer Function
*/

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
