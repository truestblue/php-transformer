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

// Use node
type Use struct {
	UseType node.Node
	Use     node.Node
	Alias   node.Node
}

// NewUse node constructor
func NewUse(UseType node.Node, use node.Node, Alias node.Node) *Use {
	return &Use{
		UseType,
		use,
		Alias,
	}
}

// Attributes returns node attributes as map
func (n *Use) Attributes() map[string]interface{} {
	return nil
}

// SetUseType set use type and returns node
func (n *Use) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Use) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Use != nil {
		vv := v.GetChildrenVisitor("Use")
		n.Use.Walk(vv)
	}

	if n.Alias != nil {
		vv := v.GetChildrenVisitor("Alias")
		n.Alias.Walk(vv)
	}

	v.LeaveNode(n)
}
