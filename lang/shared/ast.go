package shared

import "Flow2.0/lang/variables"



type Node interface {
	VisitNode() (variables.ValueNode,error)
	DisplayNode()
}