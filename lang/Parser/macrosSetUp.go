package Parser

import (
	"Flow2.0/lang/env"
	"Flow2.0/lang/shared"
)

var Macros = map[string]Macro{
	"println!": &PrintlnMacro{},
}

type Macro interface {
	Eval() (env.ValueNode, error)
	SetArgs([]shared.Node)
}
