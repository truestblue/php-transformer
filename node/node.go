package node

import "github.com/truestblue/php-transformer/walker"

// Node interface
type Node interface {
	walker.Walkable
	Attributes() map[string]interface{} // Attributes returns node attributes as map
}
