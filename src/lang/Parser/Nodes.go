package Parser

import "fmt"

// Node Main Program node /*
type Node interface {
	VisitNode() float64
	DisplayNode()
}

// NumberNode /*
type NumberNode struct {
	Value float64
}

func (n NumberNode) VisitNode() float64 {
	return n.Value
}
func (n NumberNode) DisplayNode() {
	fmt.Println(n.Value)
}

/*
Binary Operation Node
*/

type BinaryOperationNode struct {
	Left     Node
	Operator string
	Right    Node
}

func (n BinaryOperationNode) VisitNode() float64 {
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
func (n BinaryOperationNode) DisplayNode() {
	fmt.Printf("%s %s %s", n.Left, n.Operator, n.Right)
}

/*
Variable Node
*/

type VariableNode struct {
	Name  string
	Value Node
}

func (n VariableNode) VisitNode() float64 {
	return n.Value.VisitNode()
}

func (n VariableNode) DisplayNode() {
	fmt.Printf("%s = %v\n", n.Name, n.Value)
}
