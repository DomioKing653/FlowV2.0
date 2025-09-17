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
	Boolean node
*/

type BooleanNode struct {
	Value bool
}

func (n BooleanNode) VisitNode() variables.ValueNode {
	return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: n.Value}
}

func (n BooleanNode) DisplayNode() {
	fmt.Println(n.Value)
}

/*
	String node
*/

type StringNode struct {
	Value string
}

func (n StringNode) VisitNode() variables.ValueNode {
	return variables.ValueNode{Type: Lexer.StringVariable, ValueStr: n.Value}
}
func (n StringNode) DisplayNode() {
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

/*
	Loop Node
*/

type LoopNode struct {
	Nodes []Node
}

func (n LoopNode) VisitNode() variables.ValueNode {
	for {
		for _, node := range n.Nodes {
			node.VisitNode()
		}
	}
}

func (n LoopNode) DisplayNode() {

	for _, node := range n.Nodes {
		node.DisplayNode()
	}
}

/*
Compare Node
*/

type ComparisonNode struct {
	Left  Node
	Right Node
	Op    string
}

func (n ComparisonNode) VisitNode() variables.ValueNode {
	left := n.Left.VisitNode()
	right := n.Right.VisitNode()
	if left.Type == Lexer.FloatVariable && right.Type == Lexer.FloatVariable {
		switch n.Op {
		case "<":
			if left.NumberValue < right.NumberValue {
				return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: true}
			}
			return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: false}
		case ">":
			if left.NumberValue > right.NumberValue {
				return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: true}
			}
			return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: false}
		}

	}
	panic("Left or right is not float variable")
}

func (n ComparisonNode) DisplayNode() {
	fmt.Printf("{%s%s%v}\n", n.Left, n.Op, n.Right)
}

type IfNode struct {
	Expression Node
	statements []Node
}
