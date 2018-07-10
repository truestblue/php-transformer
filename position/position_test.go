<<<<<<< HEAD
package position
=======
package position_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"testing"

<<<<<<< HEAD
	"github.com/truestblue/php-transformer/position"
=======
	"github.com/z7zmey/php-parser/position"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestPrintPosition(t *testing.T) {
	pos := position.NewPosition(1, 1, 2, 5)

	expected := "Pos{Line: 1-1 Pos: 2-5}"

	actual := pos.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
