package functions

import (
    "Flow2.0/lang/Lexer"
    "Flow2.0/lang/Parser"
    "bufio"
    "fmt"
    "os"
    "strings"
)

var Console bool
var PublicCode string

func Run(code string, silent bool) {
	code = strings.TrimSpace(code)
	PublicCode = code
	if silent {
		Compile(code, true)
	} else {
		Console = true
		MainProgram(code)
	}
}

func MainProgram(code string) {
	for {
		fmt.Print(">>>")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "run" {
			Compile(code, false)
			continue
		}
		if input == "console" {
			fmt.Println("Welcome to Flow console")
			fmt.Print(">>>")
			program, _ := reader.ReadString('\n')
			Compile(program, false)
			continue
		}
		if input == "exit" {
			break
		}
		if input == "file" {
		} else {
			fmt.Println("Invalid input")
		}
	}
}

func Compile(code string, silent bool) {
	/*
	* Lexer
	 */
	lexer := Lexer.NewLexer(code)
	var tokens = lexer.Lex()

	fmt.Println("Output Tokens:")
	if !silent {
		for _, token := range tokens {
			fmt.Println(token)
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
