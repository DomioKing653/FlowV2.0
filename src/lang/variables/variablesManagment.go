package variables

import "Flow2.0/src/lang/Lexer"

var Variables map[string]*Variable

type Variable struct {
	Value    ValueNode
	Type     Lexer.VariableType
	Constant bool
}

func Init() {
	Variables = make(map[string]*Variable)
}
