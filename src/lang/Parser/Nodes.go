package Parser

// Node Main Program node /*
type Node interface {
	VisitNode() int
}

// NumberNode /*
type NumberNode struct {
	Value int
}

func (n NumberNode) VisitNode() int {
	return n.Value
}

// BinaryOperationNode /*
type BinaryOperationNode struct {
	Left     Node
	Operator string
	Right    Node
}

func (n BinaryOperationNode) VisitNode() int {
	switch n.Operator {
	case "+":
		return n.Left.VisitNode() + n.Right.VisitNode()
	case "-":
		return n.Left.VisitNode() - n.Right.VisitNode()
	case "*":
		return n.Left.VisitNode() * n.Right.VisitNode()
	case "/":
		return n.Left.VisitNode() / n.Right.VisitNode()
	default:
		panic("Idk")
	}
}
