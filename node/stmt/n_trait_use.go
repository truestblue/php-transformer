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

// TraitUse node
type TraitUse struct {
	Traits              []node.Node
	TraitAdaptationList *TraitAdaptationList
}

// NewTraitUse node constructor
func NewTraitUse(Traits []node.Node, InnerAdaptationList *TraitAdaptationList) *TraitUse {
	return &TraitUse{
		Traits,
		InnerAdaptationList,
	}
}

// Attributes returns node attributes as map
func (n *TraitUse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Traits != nil {
		vv := v.GetChildrenVisitor("Traits")
		for _, nn := range n.Traits {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.TraitAdaptationList != nil {
		vv := v.GetChildrenVisitor("TraitAdaptationList")
		n.TraitAdaptationList.Walk(vv)
	}

	v.LeaveNode(n)
}
