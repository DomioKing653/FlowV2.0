package Parser

import "fmt"

// Node Main Program node /*
type Node interface {
	VisitNode() float64
	DisplayNode()
}

/*
Number Node
*/

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
	panic("Not implemented(Variable Declaration)")
}

func (n VariableNode) DisplayNode() {
	fmt.Printf("let %s = %v\n", n.Name, n.Value)
}

/*
Variable Access Node
*/

type VariableAccessNode struct {
	Name string
}

func (n VariableAccessNode) VisitNode() float64 {
	panic("Not implemented(Access Declaration)")
}
func (n VariableAccessNode) DisplayNode() {
	fmt.Printf("%s\n", n.Name)
}

/*
	Variable Set Node
*/

type VariableAssignNode struct {
	Name  string
	Value Node
}

func (n VariableAssignNode) VisitNode() float64 {
	panic("Not implemented(Set Declaration)")
}
func (n VariableAssignNode) DisplayNode() {
	fmt.Printf("%s=%s\n", n.Name, n.Value)
}
