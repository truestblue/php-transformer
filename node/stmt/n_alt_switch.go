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

// AltSwitch node
type AltSwitch struct {
	Cond     node.Node
	CaseList *CaseList
}

// NewAltSwitch node constructor
func NewAltSwitch(Cond node.Node, CaseList *CaseList) *AltSwitch {
	return &AltSwitch{
		Cond,
		CaseList,
	}
}

// Attributes returns node attributes as map
func (n *AltSwitch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltSwitch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.CaseList != nil {
		vv := v.GetChildrenVisitor("CaseList")
		n.CaseList.Walk(vv)
	}

	v.LeaveNode(n)
}
