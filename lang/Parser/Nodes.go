package Parser

import (
	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/shared"
	"Flow2.0/lang/variables"
	"errors"
	"fmt"
)

/*
	Number Node
*/

type NumberNode struct {
	Value float64
}

func (n NumberNode) VisitNode() (variables.ValueNode,error) {
	return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: n.Value},nil
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

func (n BooleanNode) VisitNode() (variables.ValueNode,error) {
	return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: n.Value},nil
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

func (n StringNode) VisitNode() (variables.ValueNode,error) {
	return variables.ValueNode{Type: Lexer.StringVariable, ValueStr: n.Value},nil
}
func (n StringNode) DisplayNode() {
	fmt.Println(n.Value)
}

/*
	Binary Operation Node
*/

type BinaryOperationNode struct {
	Left     shared.Node
	Operator string
	Right    shared.Node
}

func (n BinaryOperationNode) VisitNode() (variables.ValueNode,error) {
	var leftValue float64 = 0
	var rightValue float64 = 0
	lftVal,err:=n.Left.VisitNode()
	CheckRuntimeErr(err)
	if lftVal.Type == Lexer.FloatVariable {
		leftValue = lftVal.NumberValue
		rghtVal,err:=n.Right.VisitNode()
		shared.Check(err)
		if rghtVal.Type == Lexer.FloatVariable {
			rightValue = rghtVal.NumberValue
		}
	} else {
		return variables.ValueNode{},errors.New("LEFT or RIGHT isnt number")
	}
	switch n.Operator {
	case "+":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue + rightValue, ValueStr: "", ValueBool: false},nil
	case "-":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue - rightValue, ValueStr: "", ValueBool: false},nil
	case "*":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue * rightValue, ValueStr: "", ValueBool: false},nil
	case "/":
		return variables.ValueNode{Type: Lexer.FloatVariable, NumberValue: leftValue / rightValue, ValueStr: "", ValueBool: false},nil
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
	statements []shared.Node
}

func (n ProgramNode) VisitNode() (variables.ValueNode,error) {
	variables.Init()
	for _, statement := range n.statements {
		_,err:=statement.VisitNode()
		CheckRuntimeErr(err)
	}
	return variables.ValueNode{},nil
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
	Value    shared.Node
	Constant bool
}

func (n VariableNode) VisitNode() (variables.ValueNode,error) {
	if _, ok := variables.Variables[n.Name]; ok {
		panic("Variable already exists")
	}
	if _, ok := variables.Variables[n.Name]; ok {
		panic("Variable already exists")
	}

	value,err := n.Value.VisitNode()
	CheckRuntimeErr(err)

	variables.Variables[n.Name] = &variables.Variable{
		Value:    value,
		Type:     value.Type,
		Constant: n.Constant,
	}
	val,err:=n.Value.VisitNode()
	CheckRuntimeErr(err)
	return val,nil
}

func (n VariableNode) DisplayNode() {
	val,err:=n.Value.VisitNode()
	CheckRuntimeErr(err)
	switch val.Type {
	case Lexer.FloatVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, val.NumberValue)
	case Lexer.BooleanVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, val.ValueBool)
	case Lexer.StringVariable:
		fmt.Printf("{let %s = %v}\n", n.Name, val.ValueStr)
	}

}

/*
	Variable Access Node
*/

type VariableAccessNode struct {
	Name string
}

func (n VariableAccessNode) VisitNode() (variables.ValueNode,error) {
	variable, ok := variables.Variables[n.Name]
	if ok {
		return variable.Value,nil
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
	Value shared.Node
}

func (n VariableAssignNode) VisitNode() (variables.ValueNode,error) {
	variable, ok := variables.Variables[n.Name]
	if !ok {
		return variables.ValueNode{},errors.New("variable not found")

	}
	if variable.Constant == true {
		return variables.ValueNode{},errors.New("variable is constant")
	}
	value,err := n.Value.VisitNode()
	CheckRuntimeErr(err)
	variable.Value = value
	return value,nil
}
func (n VariableAssignNode) DisplayNode() {
	fmt.Printf("{%s=%v}\n", n.Name, n.Value)
}

/*
	Println node
*/

type PrintLnNode struct {
	Value shared.Node
}

func (n PrintLnNode) VisitNode() (variables.ValueNode,error) {
	value,err := n.Value.VisitNode()
	CheckRuntimeErr(err)
	switch value.Type {
	case Lexer.StringVariable:
		fmt.Println(value.ValueStr)
	case Lexer.BooleanVariable:
		fmt.Println(value.ValueBool)
	case Lexer.FloatVariable:
		fmt.Println(value.NumberValue)
	}

	return value,nil
}

func (n PrintLnNode) DisplayNode() {
	fmt.Printf("prinln()\n")
}

/*
	Loop Node
*/

type LoopNode struct {
	Nodes []shared.Node
}

func (n LoopNode) VisitNode() (variables.ValueNode,error) {
	for {
		for _, node := range n.Nodes {
			_,err :=node.VisitNode()
			CheckRuntimeErr(err)
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
	Left  shared.Node
	Right shared.Node
	Op    string
}

func (n ComparisonNode) VisitNode() (variables.ValueNode,error) {
	left,err := n.Left.VisitNode()
	CheckRuntimeErr(err)
	right,err := n.Right.VisitNode()
	CheckRuntimeErr(err)
	if left.Type == Lexer.FloatVariable && right.Type == Lexer.FloatVariable {
		switch n.Op {
		case "<":
			if left.NumberValue < right.NumberValue {
				return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: true},nil
			}
			return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: false},nil
		case ">":
			if left.NumberValue > right.NumberValue {
				return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: true},nil
			}
			return variables.ValueNode{Type: Lexer.BooleanVariable, ValueBool: false},nil
		}

	}
	return variables.ValueNode{},errors.New("left or right isnt float")
}

func (n ComparisonNode) DisplayNode() {
	fmt.Printf("{%s%s%v}\n", n.Left, n.Op, n.Right)
}

/*
 	If node
*/

type IfNode struct {
	Expression shared.Node
	statements []shared.Node
}

func (n IfNode) VisitNode() (variables.ValueNode,error) {
	expr,err:=n.Expression.VisitNode()
	CheckRuntimeErr(err)
	if expr.Type==Lexer.BooleanVariable{
		expr,err=n.Expression.VisitNode()
		CheckRuntimeErr(err)
		if expr.ValueBool{
			for _,statment:=range n.statements{
				_,err:=statment.VisitNode()
				CheckRuntimeErr(err)
			}
		}
	}
	return variables.ValueNode{},nil
}
func (n IfNode) DisplayNode() {
	fmt.Print("if(")
	n.Expression.DisplayNode()
	fmt.Print(")")
	for _, node := range n.statements {
		node.DisplayNode()
	}
}
/*
	while node
*/

type WhileNpde struct {
	Expression shared.Node
	Statments  []shared.Node
}

func (n WhileNpde) VisitNode() (variables.ValueNode,error)  {
	expr,err:=n.Expression.VisitNode()
	CheckRuntimeErr(err)
	if expr.Type==Lexer.BooleanVariable{
		for{
			if !expr.ValueBool{
				break
			}else{
				for _,statment:= range n.Statments{
					_,err:=statment.VisitNode()
					CheckRuntimeErr(err)
				}
			}
		}
	}
	return expr,nil
}