package stmt

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Break node
type Break struct {
	Expr node.Node
}

// NewBreak node constructor
func NewBreak(Expr node.Node) *Break {
	return &Break{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Break) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Break) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
