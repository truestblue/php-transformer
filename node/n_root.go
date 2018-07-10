package node

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/walker"
=======
	"github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Root node
type Root struct {
	Stmts []Node
}

// NewRoot node constructor
func NewRoot(Stmts []Node) *Root {
	return &Root{
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *Root) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Root) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
