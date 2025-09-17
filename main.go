package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Flow2.0/src/functions"
	"Flow2.0/src/lang/Lexer"
	"Flow2.0/src/lang/Parser"
)

var silent = false

func mainProgram(code string) {
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

var programPath string

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
func Run() {
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
	if silent {
		Compile(code, true)
	} else {
		mainProgram(code)
	}
}

func main() {
	Run()
}
