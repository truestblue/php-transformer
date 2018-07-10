package position

import (
	"testing"

	"github.com/truestblue/php-transformer/position"
)

func TestPrintPosition(t *testing.T) {
	pos := position.NewPosition(1, 1, 2, 5)

	expected := "Pos{Line: 1-1 Pos: 2-5}"

	actual := pos.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
