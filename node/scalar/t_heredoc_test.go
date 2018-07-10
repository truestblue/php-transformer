<<<<<<< HEAD
package scalar
=======
package scalar_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"bytes"
	"testing"

<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node/expr"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
=======
	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestHeredocSimpleLabel(t *testing.T) {
	src := `<? <<<LBL
test $var
LBL;
`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "LBL",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "var"}},
						&scalar.EncapsedStringPart{Value: "\n"},
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

func TestSimpleHeredocLabel(t *testing.T) {
	src := `<? <<<"LBL"
test $var
LBL;
`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "\"LBL\"",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "var"}},
						&scalar.EncapsedStringPart{Value: "\n"},
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

func TestSimpleNowdocLabel(t *testing.T) {
	src := `<? <<<'LBL'
test $var
LBL;
`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "'LBL'",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test $var\n"},
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

func TestEmptyHeredoc(t *testing.T) {
	src := `<? <<<CAD
CAD;
`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "CAD",
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

func TestHeredocScalarString(t *testing.T) {
	src := `<? <<<CAD
	hello
CAD;
`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Heredoc{
					Label: "CAD",
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "\thello\n"},
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
