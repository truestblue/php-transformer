package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-transformer/node/expr"

	"github.com/z7zmey/php-transformer/node/scalar"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestSimpleEcho(t *testing.T) {
	src := `<? echo $a, 1;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
					&scalar.Lnumber{Value: "1"},
				},
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

func TestEcho(t *testing.T) {
	src := `<? echo($a);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "a"},
					},
				},
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
