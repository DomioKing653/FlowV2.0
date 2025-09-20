package shared

import "Flow2.0/lang/env"

type Node interface {
	VisitNode() (env.ValueNode, error)
	DisplayNode()
}
