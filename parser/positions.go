package parser

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/position"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Positions a collection of positions attached to nodes
type Positions map[node.Node]*position.Position

// AddPosition attaches a position to a node
func (p Positions) AddPosition(node node.Node, position *position.Position) {
	p[node] = position
}
