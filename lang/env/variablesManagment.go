package env

import "Flow2.0/lang/Lexer"

var Scopes []Scope

type Variable struct {
	Value    ValueNode
	Type     Lexer.VariableType
	Constant bool
}
