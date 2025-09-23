package Parser

import (
	"fmt"

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
