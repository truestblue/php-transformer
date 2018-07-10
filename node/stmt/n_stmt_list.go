package stmt

import (
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
)

// StmtList node
type StmtList struct {
	Stmts []node.Node
}

// NewStmtList node constructor
func NewStmtList(Stmts []node.Node) *StmtList {
	return &StmtList{
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *StmtList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StmtList) Walk(v walker.Visitor) {
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
