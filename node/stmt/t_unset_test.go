package stmt

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/expr"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestUnset(t *testing.T) {
	src := `<? unset($a);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestUnsetVars(t *testing.T) {
	src := `<? unset($a, $b);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestUnsetTrailingComma(t *testing.T) {
	src := `<? unset($a, $b,);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Unset{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
