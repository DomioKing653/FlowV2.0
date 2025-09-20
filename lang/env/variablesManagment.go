package env

import "Flow2.0/lang/Lexer"

var Variables map[string]*Variable

type Variable struct {
	Value    ValueNode
	Type     Lexer.VariableType
	Constant bool
}

func Init() {
	Variables = make(map[string]*Variable)
}
