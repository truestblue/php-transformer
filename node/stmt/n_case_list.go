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

// CaseList node
type CaseList struct {
	Cases []node.Node
}

// NewCaseList node constructor
func NewCaseList(Cases []node.Node) *CaseList {
	return &CaseList{
		Cases,
	}
}

// Attributes returns node attributes as map
func (n *CaseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CaseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cases != nil {
		vv := v.GetChildrenVisitor("Cases")
		for _, nn := range n.Cases {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
