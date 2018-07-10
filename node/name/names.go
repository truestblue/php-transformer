package name

import (
	"github.com/truestblue/php-transformer/node"
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
