package name

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// FullyQualified node
type FullyQualified struct {
	Parts []node.Node
}

// NewFullyQualified node constructor
func NewFullyQualified(Parts []node.Node) *FullyQualified {
	return &FullyQualified{
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *FullyQualified) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *FullyQualified) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}

// GetParts returns the name parts
func (n *FullyQualified) GetParts() []node.Node {
	return n.Parts
}
