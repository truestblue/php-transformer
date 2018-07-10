package expr

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/walker"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// StaticCall node
type StaticCall struct {
	Class        node.Node
	Call         node.Node
	ArgumentList *node.ArgumentList
}

// NewStaticCall node constructor
func NewStaticCall(Class node.Node, Call node.Node, ArgumentList *node.ArgumentList) *StaticCall {
	return &StaticCall{
		Class,
		Call,
		ArgumentList,
	}
}

// Attributes returns node attributes as map
func (n *StaticCall) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticCall) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.Call != nil {
		vv := v.GetChildrenVisitor("Call")
		n.Call.Walk(vv)
	}

	if n.ArgumentList != nil {
		vv := v.GetChildrenVisitor("ArgumentList")
		n.ArgumentList.Walk(vv)
	}

	v.LeaveNode(n)
}
