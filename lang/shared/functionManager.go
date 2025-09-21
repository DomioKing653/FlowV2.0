package shared

var Functions map[string]*Function

type Function struct {
	Nodes []Node
	Args  []string
}

func InitFunctions() {
	Functions = make(map[string]*Function)
}
