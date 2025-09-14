package main

import (
	"fmt"
	"os"
	"strings"

	"Flow2.0/functions"
	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/Parser"
)

func mainProgram(code string) {
	/*
	* Lexer
	 */
	lexer := Lexer.NewLexer(code)
	var tokens []Lexer.Token = lexer.Lex()
	fmt.Println("Output Tokens:")
	for _, tok := range tokens {
		fmt.Printf("Token: %-6s  Value: %s\n", tok.Type, tok.Value)
	}
	/*
	* Parser
	 */
	parser := Parser.NewParser(tokens)
	ast := parser.Parse()
	ast.DisplayNode()
	/*
	 * Interpreter
	 */
	fmt.Println(ast.VisitNode())
}
func main() {
	programPath := os.Args[1]
	code := functions.ReadFile(programPath)
	code = strings.TrimSpace(code)
	mainProgram(code)
}
