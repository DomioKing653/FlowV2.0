package Parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/env"
	"Flow2.0/lang/shared"
)

/*
Println macro
*/
type PrintlnMacro struct {
	args []shared.Node
}

func (f *PrintlnMacro) Eval() (env.ValueNode, error) {
	for _, n := range f.args {
		val, err := n.VisitNode()
		CheckRuntimeErr(err)
		switch val.Type {
		case Lexer.BooleanVariable:
			fmt.Println(val.ValueBool)
		case Lexer.FloatVariable:
			fmt.Println(val.NumberValue)
		case Lexer.StringVariable:
			fmt.Println(val.ValueStr)
		}
	}
	return env.ValueNode{}, nil
}

func (f *PrintlnMacro) SetArgs(args []shared.Node) {
	f.args = args
}

/*
ReadFloat Macro
*/

type ReadFloatMacro struct {
	args []shared.Node
}

func (f *ReadFloatMacro) Eval() (env.ValueNode, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter float:")
	line, err := reader.ReadString('\n')
	CheckRuntimeErr(err)
	value, err := strconv.ParseFloat(line, 64)
	return env.ValueNode{Type: Lexer.FloatVariable, NumberValue: value}, nil
}

func (f *ReadFloatMacro) SetArgs(args []shared.Node) {
	f.args = args
}
