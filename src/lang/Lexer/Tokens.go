package Lexer

type TokenType string

const (
	NUMBER TokenType = "NUMBER"
	PLUS   TokenType = "PLUS"
	MINUS  TokenType = "MINUS"
	STAR   TokenType = "STAR"
	SLASH  TokenType = "SLASH"
	EOF    TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}
