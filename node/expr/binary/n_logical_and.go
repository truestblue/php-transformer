package binary

import (
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
)

// LogicalAnd node
type LogicalAnd struct {
	Left  node.Node
	Right node.Node
}

// NewLogicalAnd node constructor
func NewLogicalAnd(Variable node.Node, Expression node.Node) *LogicalAnd {
	return &LogicalAnd{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *LogicalAnd) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *LogicalAnd) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
