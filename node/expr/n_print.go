package expr

import (
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
)

// Print node
type Print struct {
	Expr node.Node
}

// NewPrint node constructor
func NewPrint(Expression node.Node) *Print {
	return &Print{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Print) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Print) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
