package main

import (
	"fmt"
	"os"
	"strings"

	"Flow2.0/functions"
	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/Parser"
)

var silent = false

func mainProgram(code string) {
	/*
	* Lexer
	 */
	lexer := Lexer.NewLexer(code)
	var tokens = lexer.Lex()

	if !silent {
		fmt.Println("Output Tokens:")
		for _, tok := range tokens {
			fmt.Printf("Token: %-6s  Value: %s\n", tok.Type, tok.Value)
		}
	}

	/*
	* Parser
	 */
	parser := Parser.NewParser(tokens)
	ast := parser.Parse()
	if !silent {
		ast.DisplayNode()
	}
	/*
	 * Interpreter
	 */
	ast.VisitNode()
}

var programPath string

func main() {
	if len(os.Args) > 1 {
		programPath = os.Args[1]
		if len(os.Args) > 2 {
			if os.Args[2] == "-s" {
				silent = true
			}
		}
	}
	code := functions.ReadFile(programPath)
	code = strings.TrimSpace(code)
	mainProgram(code)
}
