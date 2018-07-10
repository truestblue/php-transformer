package scanner

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-transformer/position"

	"github.com/z7zmey/php-transformer/comment"

	"github.com/z7zmey/php-transformer/scanner"
)

func TestToken(t *testing.T) {
	pos := position.NewPosition(1, 1, 0, 3)
	tkn := &scanner.Token{
		Value:    `foo`,
		Position: pos,
	}

	c := []*comment.Comment{
		comment.NewComment("test comment", nil),
	}

	tkn.Comments = c

	if !reflect.DeepEqual(tkn.Comments, c) {
		t.Errorf("comments are not equal\n")
	}

	if tkn.String() != `foo` {
		t.Errorf("token value is not equal\n")
	}

	if tkn.Position != pos {
		t.Errorf("token position is not equal\n")
	}
}
