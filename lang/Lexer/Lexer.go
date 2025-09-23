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
		case '#':
			l.advance()
			for {
				if l.peek() == '\n' {
					break
				}
				l.advance()
			}
		case '{':
			l.advance()
			return Token{Type: OpeningParen, Value: "{"}
		case '}':
			l.advance()
			return Token{Type: ClosingParen, Value: "}"}
		case '>':
			l.advance()
			return Token{Type: GREATER, Value: ">"}
		case '<':
			l.advance()
			return Token{Type: LESS, Value: "<"}
		case '"':
			l.advance()
			return l.MakeString()
		case ',':
			l.advance()
			return Token{Type: COMMA, Value: ","}
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
		} else if l.peek() == '.' && !hasDot {
			hasDot = true
			l.advance()
		} else if l.peek() == '.' && hasDot {
			panic("can't have more than one dot in a number")
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
	for unicode.IsLetter(l.peek()) || l.peek() == '!' {
		l.advance()
	}
	text := l.input[start:l.pos]
	switch text {
	case "let":
		return Token{Type: LET, Value: "let"}
	case "println":
		return Token{Type: PRINTLN, Value: "println"}
	case "const":
		return Token{Type: CONST, Value: "const"}
	case "loop":
		return Token{Type: LOOP, Value: "loop"}
	case "true":
		return Token{Type: BOOL, Value: "true"}
	case "false":
		return Token{Type: BOOL, Value: "false"}
	case "fn":
		return Token{Type: FUNCTION, Value: "fn"}
	case "if":
		return Token{Type: IF, Value: "if"}
	case "while":
		return Token{Type: WHILE, Value: "while"}
	case "break":
		return Token{Type: BREAK, Value: "break"}
	case "return":
		return Token{Type: RETURN, Value: "return"}
	default:
		return Token{Type: IDENTIFIER, Value: text}
	}
}
func (l *Lexer) MakeString() Token {
	start := l.pos
	for {
		ch := l.peek()
		if ch == 0 {
			panic("unterminated string literal")
		}
		if ch == '"' {
			break
		}
		l.advance()
	}
	text := l.input[start:l.pos]
	l.advance()
	return Token{Type: STRING, Value: text}
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
	fmt.Printf("Lexing ended")
	return tokens
}
