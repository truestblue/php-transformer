package parser

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/errors"
	"github.com/truestblue/php-transformer/node"
=======
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Parser interface
type Parser interface {
	Parse() int
	GetPath() string
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	GetComments() Comments
	GetPositions() Positions
}
