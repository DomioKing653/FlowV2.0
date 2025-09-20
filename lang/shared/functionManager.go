package shared

var Functions map[string]*Function

type Function struct {
	Nodes []Node
	Args  []string
}
