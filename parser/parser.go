package parser

import (
	"github.com/truestblue/php-transformer/errors"
	"github.com/truestblue/php-transformer/node"
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
