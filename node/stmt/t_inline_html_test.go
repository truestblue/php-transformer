package stmt_test

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestInlineHtml(t *testing.T) {
	src := `<? ?> <div></div>`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Nop{},
			&stmt.InlineHtml{Value: "<div></div>"},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
