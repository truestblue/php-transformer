package stmt_test

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/expr"
	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestStaticVar(t *testing.T) {
	src := `<? static $a;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestStaticVars(t *testing.T) {
	src := `<? static $a, $b = 1;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Expr:     &scalar.Lnumber{Value: "1"},
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

func TestStaticVars2(t *testing.T) {
	src := `<? static $a = 1, $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &scalar.Lnumber{Value: "1"},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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
