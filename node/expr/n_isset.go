package expr

import (
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
)

// Isset node
type Isset struct {
	Variables []node.Node
}

// NewIsset node constructor
func NewIsset(Variables []node.Node) *Isset {
	return &Isset{
		Variables,
	}
}

// Attributes returns node attributes as map
func (n *Isset) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Isset) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variables != nil {
		vv := v.GetChildrenVisitor("Variables")
		for _, nn := range n.Variables {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
