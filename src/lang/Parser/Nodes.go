package Parser

import (
	"fmt"

	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/variables"
)

// Node Main Program node /*
type Node interface {
	VisitNode() variables.ValueNode
	DisplayNode()
}

/*
	Number Node
*/

type NumberNode struct {
	Value float64
}

func (n NumberNode) VisitNode() variables.ValueNode {
	return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: n.Value}
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

func (n BinaryOperationNode) VisitNode() variables.ValueNode {
	var leftValue float64 = 0
	var rightValue float64 = 0
	if n.Left.VisitNode().Type == Lexer.FloatVariable {
		leftValue = n.Left.VisitNode().NumberValue
		if n.Right.VisitNode().Type == Lexer.FloatVariable {
			rightValue = n.Right.VisitNode().NumberValue
		}
	} else {
		panic("Left is not a float value")
	}
	switch n.Operator {
	case "+":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue + rightValue, ValueStr: "", ValueBool: false}
	case "-":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue - rightValue, ValueStr: "", ValueBool: false}
	case "*":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue * rightValue, ValueStr: "", ValueBool: false}
	case "/":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue / rightValue, ValueStr: "", ValueBool: false}
	default:
		panic("Unknown operator")
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

func (n ProgramNode) VisitNode() variables.ValueNode {
	variables.Init()
	for _, statement := range n.statements {
		statement.VisitNode()
	}
	return variables.ValueNode{}
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
	Name     string
	Value    Node
	Constant bool
}

func (n VariableNode) VisitNode() variables.ValueNode {
	if _, ok := variables.Variables[n.Name]; ok {
		panic("Variable already exists")
	}
	if _, ok := variables.Variables[n.Name]; ok {
		panic("Variable already exists")
	}

	value := n.Value.VisitNode()

	variables.Variables[n.Name] = &variables.Variable{
		Value:    value,
		Type:     value.Type,
		Constant: n.Constant,
	}

	return value

	return n.Value.VisitNode()
}

func (n VariableNode) DisplayNode() {
	switch n.Value.VisitNode().Type {
	case Lexer.FloatVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, n.Value.VisitNode().NumberValue)
	case Lexer.BooleanVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, n.Value.VisitNode().ValueBool)
	case Lexer.StringVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, n.Value.VisitNode().ValueStr)
	}

}

/*
	Variable Access Node
*/

type VariableAccessNode struct {
	Name string
}

func (n VariableAccessNode) VisitNode() variables.ValueNode {
	variable, ok := variables.Variables[n.Name]
	if ok {
		return variable.Value
	} else {
		panic(fmt.Sprintf("Variable %s not found", n.Name))
	}
	return variables.Variables[n.Name].Value
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

func (n VariableAssignNode) VisitNode() variables.ValueNode {
	variable, ok := variables.Variables[n.Name]
	if !ok {
		panic(fmt.Sprintf("Variable %s not found", n.Name))
	}
	if variable.Constant == true {
		panic(fmt.Sprintf("Variable %s is constant", n.Name))
	}
	value := n.Value.VisitNode()
	variable.Value = value
	return value
}
func (n VariableAssignNode) DisplayNode() {
	fmt.Printf("{%s=%v}\n", n.Name, n.Value)
}

/*
	Println node
*/

type PrintLnNode struct {
	Value Node
}

func (n PrintLnNode) VisitNode() variables.ValueNode {
	value := n.Value.VisitNode()
	switch value.Type {
	case Lexer.StringVariable:
		fmt.Println(value.ValueStr)
	case Lexer.BooleanVariable:
		fmt.Println(value.ValueBool)
	case Lexer.FloatVariable:
		fmt.Println(value.NumberValue)
	}

	return value
}

func (n PrintLnNode) DisplayNode() {
	fmt.Printf("prinln(%v)\n", n.Value)
}
