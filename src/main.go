package main

import (
	"fmt"
	"os"
	"strings"

	"Flow2.0/functions"
	"Flow2.0/lang"
)

func main() {
	programPath := os.Args[1]
	code := functions.ReadFile(programPath)
	code = strings.TrimSpace(code)
	lexer := lang.NewLexer(code)
	tokens := lexer.Lex()

	fmt.Println("Output Tokens:")
	for _, tok := range tokens {
		fmt.Printf("Token: %-6s  Value: %s\n", tok.Type, tok.Value)
	}
}
