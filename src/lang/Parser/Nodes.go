package Parser

import (
	"fmt"

	"Flow2.0/lang/variables"
)

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
	fmt.Printf("{%s %s %s}", n.Left, n.Operator, n.Right)
}

/*
Program Node
*/

type ProgramNode struct {
	statements []Node
}

func (n ProgramNode) VisitNode() float64 {
	for _, statement := range n.statements {
		statement.VisitNode()
	}
	return 0
}
func (n ProgramNode) DisplayNode() {
	fmt.Printf("Program{\n")
	for _, statement := range n.statements {
		fmt.Print("  ")
		statement.DisplayNode()
	}
	fmt.Printf("}\n")

}

/*
Variable Node
*/

type VariableNode struct {
	Name  string
	Value Node
}

func (n VariableNode) VisitNode() float64 {
	_, ok := variables.Variables[n.Name]
	if ok {
		panic("Variable already exists")
	} else {
		variables.Variables[n.Name] = n.Value.VisitNode()
	}
	return n.Value.VisitNode()
}

func (n VariableNode) DisplayNode() {
	fmt.Printf("{let %s = %v}\n", n.Name, n.Value.VisitNode())
}

/*
Variable Access Node
*/

type VariableAccessNode struct {
	Name string
}

func (n VariableAccessNode) VisitNode() float64 {
	variable, ok := variables.Variables[n.Name]
	if ok {
		return variable
	} else {
		panic(fmt.Sprintf("Variable %s not found", n.Name))
	}
	return variables.Variables[n.Name]
}
func (n VariableAccessNode) DisplayNode() {
	fmt.Printf("{%s}\n", n.Name)
}

/*
	Variable Set Node
*/

type VariableAssignNode struct {
	Name  string
	Value Node
}

func (n VariableAssignNode) VisitNode() float64 {
	_, ok := variables.Variables[n.Name]
	if ok {
		variables.Variables[n.Name] = n.Value.VisitNode()
		return 0
	} else {
		panic(fmt.Sprintf("Variable %s not found", n.Name))
	}
}
func (n VariableAssignNode) DisplayNode() {
	fmt.Printf("{%s=%v}\n", n.Name, n.Value)
}
