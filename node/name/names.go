package name

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
=======
	"github.com/z7zmey/php-parser/node"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
