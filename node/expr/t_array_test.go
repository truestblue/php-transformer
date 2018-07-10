package expr

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node/scalar"

	"github.com/truestblue/php-transformer/node/expr"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestArray(t *testing.T) {
	src := `<? array();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{},
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

func TestArrayItem(t *testing.T) {
	src := `<? array(1);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Val: &scalar.Lnumber{Value: "1"},
						},
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

func TestArrayItems(t *testing.T) {
	src := `<? array(1=>1, &$b,);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							Key: &scalar.Lnumber{Value: "1"},
							Val: &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							Val: &expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}}},
						},
						nil,
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
