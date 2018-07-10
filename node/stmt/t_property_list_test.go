package stmt

import (
	"bytes"
	"github.com/z7zmey/php-transformer/node/expr"
	"github.com/z7zmey/php-transformer/node/scalar"
	"testing"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestProperty(t *testing.T) {
	src := `<? class foo {var $a;}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "var"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
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

func TestProperties(t *testing.T) {
	src := `<? class foo {public static $a, $b = 1;}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
							&node.Identifier{Value: "static"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "b"}},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
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

func TestProperties2(t *testing.T) {
	src := `<? class foo {public static $a = 1, $b;}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.PropertyList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
							&node.Identifier{Value: "static"},
						},
						Properties: []node.Node{
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Property{
								PhpDocComment: "",
								Variable:      &expr.Variable{VarName: &node.Identifier{Value: "b"}},
							},
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
