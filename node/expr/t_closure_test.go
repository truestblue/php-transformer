package expr_test

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node/name"

	"github.com/truestblue/php-transformer/node/expr"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestClosure(t *testing.T) {
	src := `<? function(){};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Stmts:         []node.Node{},
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

func TestClosureUse(t *testing.T) {
	src := `<? function($a, $b) use ($c, &$d) {};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						},
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Uses: []node.Node{
							&expr.Variable{VarName: &node.Identifier{Value: "c"}},
							&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "d"}}},
						},
					},
					Stmts: []node.Node{},
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

func TestClosureUse2(t *testing.T) {
	src := `<? function($a, $b) use (&$c, $d) {};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					Params: []node.Node{
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						},
						&node.Parameter{
							ByRef:    false,
							Variadic: false,
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						},
					},
					ClosureUse: &expr.ClosureUse{
						Uses: []node.Node{
							&expr.Reference{Variable: &expr.Variable{VarName: &node.Identifier{Value: "c"}}},
							&expr.Variable{VarName: &node.Identifier{Value: "d"}},
						},
					},
					Stmts: []node.Node{},
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

func TestClosureReturnType(t *testing.T) {
	src := `<? function(): void {};`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Closure{
					ReturnsRef:    false,
					Static:        false,
					PhpDocComment: "",
					ReturnType: &name.Name{
						Parts: []node.Node{&name.NamePart{Value: "void"}},
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
