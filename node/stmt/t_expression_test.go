package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-transformer/node/scalar"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestExpression(t *testing.T) {
	src := `<? 1;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "1"},
			},
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
