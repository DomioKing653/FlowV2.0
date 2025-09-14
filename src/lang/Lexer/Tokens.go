package Lexer

type TokenType string

const (
	/*
		Types
	*/
	INT   TokenType = "NUMBER"
	FLOAT TokenType = "FLOAT"
	/*
		Keywords
	*/
	LET TokenType = "LET"

	/*
		Functions
	*/
	PRINTLN TokenType = "PRINTLN"
	/*
		Misc Thingy
	*/

	IDENTIFIER TokenType = "IDENTIFIER"
	EQUALS     TokenType = "EQUALS"
	/*
		Math
	*/
	PLUS   TokenType = "PLUS"
	MINUS  TokenType = "MINUS"
	STAR   TokenType = "STAR"
	SLASH  TokenType = "SLASH"
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	/*
		Parser Thingy
	*/
	EOF TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}
