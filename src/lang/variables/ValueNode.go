package variables

import "Flow2.0/src/lang/Lexer"

type ValueNode struct {
	Type        Lexer.VariableType
	NumberValue float64
	ValueStr    string
	ValueBool   bool
}
