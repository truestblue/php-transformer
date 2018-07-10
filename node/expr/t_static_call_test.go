package expr

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-transformer/node/name"

	"github.com/z7zmey/php-transformer/node/expr"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestStaticCall(t *testing.T) {
	src := `<? Foo::bar();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
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

func TestStaticCallRelative(t *testing.T) {
	src := `<? namespace\Foo::bar();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
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

func TestStaticCallFullyQualified(t *testing.T) {
	src := `<? \Foo::bar();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &node.Identifier{Value: "bar"},
					ArgumentList: &node.ArgumentList{},
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

func TestStaticCallVar(t *testing.T) {
	src := `<? Foo::$bar();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
					Call:         &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					ArgumentList: &node.ArgumentList{},
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

func TestStaticCallVarVar(t *testing.T) {
	src := `<? $foo::$bar();`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.StaticCall{
					Class:        &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
					Call:         &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
					ArgumentList: &node.ArgumentList{},
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
