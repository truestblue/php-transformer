package parser

import (
	"testing"

	"github.com/truestblue/php-transformer/comment"
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/parser"
)

func TestComments(t *testing.T) {
	n := node.NewIdentifier("test")

	commentGroup := []*comment.Comment{
		comment.NewComment("/** hello world */", nil),
		comment.NewComment("// hello world", nil),
	}

	comments := parser.Comments{}
	comments.AddComments(n, commentGroup)

	if comments[n][0].String() != "/** hello world */" {
		t.Errorf("expected and actual are not equal\n")
	}
	if comments[n][1].String() != "// hello world" {
		t.Errorf("expected and actual are not equal\n")
	}
}
