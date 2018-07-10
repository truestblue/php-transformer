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

// Catch node
type Catch struct {
	Types    []node.Node
	Variable node.Node
	Stmts    []node.Node
}

// NewCatch node constructor
func NewCatch(Types []node.Node, Variable node.Node, Stmts []node.Node) *Catch {
	return &Catch{
		Types,
		Variable,
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *Catch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Catch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Types != nil {
		vv := v.GetChildrenVisitor("Types")
		for _, nn := range n.Types {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
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
