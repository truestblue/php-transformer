<<<<<<< HEAD
package parser
=======
package parser_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"testing"

<<<<<<< HEAD
	"github.com/truestblue/php-transformer/comment"
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/parser"
=======
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
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
