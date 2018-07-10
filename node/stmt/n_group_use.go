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

// GroupUse node
type GroupUse struct {
	UseType node.Node
	Prefix  node.Node
	UseList []node.Node
}

// NewGroupUse node constructor
func NewGroupUse(UseType node.Node, Prefix node.Node, UseList []node.Node) *GroupUse {
	return &GroupUse{
		UseType,
		Prefix,
		UseList,
	}
}

// Attributes returns node attributes as map
func (n *GroupUse) Attributes() map[string]interface{} {
	return nil
}

// SetUseType set use type and returns node
func (n *GroupUse) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *GroupUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Prefix != nil {
		vv := v.GetChildrenVisitor("Prefix")
		n.Prefix.Walk(vv)
	}

	if n.UseList != nil {
		vv := v.GetChildrenVisitor("UseList")
		for _, nn := range n.UseList {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
