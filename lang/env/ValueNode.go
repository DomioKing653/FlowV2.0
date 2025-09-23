package env

import "Flow2.0/lang/Lexer"

type ValueNode struct {
	Type        Lexer.VariableType
	NumberValue float64
	ValueStr    string
	ValueBool   bool
	Return      bool
}
