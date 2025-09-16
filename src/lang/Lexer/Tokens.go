package Lexer

type TokenType string

const (
	/*
		Types
	*/
	INT   TokenType = "NUMBER"
	FLOAT TokenType = "FLOAT"
	BOOL  TokenType = "BOOL"
	/*
		Keywords
	*/
	LET   TokenType = "LET"
	CONST TokenType = "CONST"
	LOOP  TokenType = "LOOP"

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
	PLUS         TokenType = "PLUS"
	MINUS        TokenType = "MINUS"
	STAR         TokenType = "STAR"
	SLASH        TokenType = "SLASH"
	LPAREN       TokenType = "LPAREN"
	RPAREN       TokenType = "RPAREN"
	OpeningParen TokenType = "OPENING_PAREN"
	ClosingParen TokenType = "CLOSING_PAREN"
	/*
		Parser Thingy
	*/
	EOF TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type VariableType string

const (
	FloatVariable   VariableType = "FLOAT"
	StringVariable  VariableType = "STRING"
	BooleanVariable VariableType = "BOOLEAN"
	NilVariable     VariableType = "NIL"
)

type Variable struct {
	Type  VariableType
	Value string
}
