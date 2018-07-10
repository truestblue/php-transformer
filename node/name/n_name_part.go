package name

<<<<<<< HEAD
import "github.com/truestblue/php-transformer/walker"
=======
import "github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

// NamePart node
type NamePart struct {
	Value string
}

// NewNamePart node constructor
func NewNamePart(Value string) *NamePart {
	return &NamePart{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *NamePart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *NamePart) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
