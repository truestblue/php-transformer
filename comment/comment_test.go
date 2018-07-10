package comment

import (
	"testing"

	"github.com/truestblue/php-transformer/position"

	"github.com/truestblue/php-transformer/comment"
)

func TestCommentGetPosition(t *testing.T) {
	expected := position.NewPosition(0, 0, 0, 0)

	comment := comment.NewComment("/** hello world */", expected)

	actual := comment.Position()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentPrint(t *testing.T) {
	expected := "/** hello world */"

	comment := comment.NewComment(expected, nil)

	actual := comment.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}
