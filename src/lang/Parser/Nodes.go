package Parser

import "fmt"

// Node Main Program node /*
type Node interface {
	VisitNode() string
}

// NumberNode /*
type NumberNode struct {
	Value int
}

func (n NumberNode) VisitNode() string {
	return fmt.Sprintf("%f", n.Value)
}

// BinaryOperationNode /*
type BinaryOperationNode struct {
	Left     Node
	Operator string
	Right    Node
}

func (n BinaryOperationNode) VisitNode() string {
	return fmt.Sprintf("%s %s %s", n.Left.VisitNode(), n.Operator, n.Left.VisitNode())
}
