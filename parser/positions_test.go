<<<<<<< HEAD
package parser
=======
package parser_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"testing"

<<<<<<< HEAD
	"github.com/truestblue/php-transformer/position"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/parser"
=======
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestPositions(t *testing.T) {
	n := node.NewIdentifier("test")

	expected := position.NewPosition(0, 0, 0, 0)

	positions := parser.Positions{}
	positions.AddPosition(n, expected)

	actual := positions[n]

	if actual != expected {
		t.Errorf("expected and actual are not equal\n")
	}
}
