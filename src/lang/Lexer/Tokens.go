package Lexer

type TokenType string

const (
	/*
		Types
	*/
	INT    TokenType = "NUMBER"
	FLOAT  TokenType = "FLOAT"
	STRING TokenType = "STRING"
	/*
		Keywords
	*/
	LET TokenType = "LET"

	/*
		Misc Thingy
	*/

	IDENTIFIER TokenType = "IDENTIFIER"
	EQUALS     TokenType = "EQUALS"
	/*
		Math
	*/
	PLUS  TokenType = "PLUS"
	MINUS TokenType = "MINUS"
	STAR  TokenType = "STAR"
	SLASH TokenType = "SLASH"
	/*
		Parser Thingy
	*/
	EOF TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}
