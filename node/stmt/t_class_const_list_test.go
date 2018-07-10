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

func TestClassConstList(t *testing.T) {
	src := `<? class foo{ public const FOO = 1, BAR = 2; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Modifiers: []node.Node{
							&node.Identifier{Value: "public"},
						},
						Consts: []node.Node{
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "FOO"},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "BAR"},
								Expr:          &scalar.Lnumber{Value: "2"},
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
}

func TestClassConstListWithoutModifiers(t *testing.T) {
	src := `<? class foo{ const FOO = 1, BAR = 2; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts: []node.Node{
					&stmt.ClassConstList{
						Consts: []node.Node{
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "FOO"},
								Expr:          &scalar.Lnumber{Value: "1"},
							},
							&stmt.Constant{
								PhpDocComment: "",
								ConstantName:  &node.Identifier{Value: "BAR"},
								Expr:          &scalar.Lnumber{Value: "2"},
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
