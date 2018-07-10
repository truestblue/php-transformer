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

// For node
type For struct {
	Init []node.Node
	Cond []node.Node
	Loop []node.Node
	Stmt node.Node
}

// NewFor node constructor
func NewFor(Init []node.Node, Cond []node.Node, Loop []node.Node, Stmt node.Node) *For {
	return &For{
		Init,
		Cond,
		Loop,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *For) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *For) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Init != nil {
		vv := v.GetChildrenVisitor("Init")
		for _, nn := range n.Init {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		for _, nn := range n.Cond {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Loop != nil {
		vv := v.GetChildrenVisitor("Loop")
		for _, nn := range n.Loop {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
