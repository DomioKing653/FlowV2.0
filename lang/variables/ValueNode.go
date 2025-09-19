package variables

import "Flow2.0/lang/Lexer"

type ValueNode struct {
	Breaking   	bool
	Type        Lexer.VariableType
	NumberValue float64
	ValueStr    string
	ValueBool   bool
}
