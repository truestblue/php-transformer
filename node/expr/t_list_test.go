<<<<<<< HEAD
package expr
=======
package expr_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"bytes"
	"testing"

<<<<<<< HEAD
	"github.com/z7zmey/php-transformer/node/expr/assign"

	"github.com/z7zmey/php-transformer/node/expr"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
=======
	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestEmptyList(t *testing.T) {
	src := `<? list() = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestList(t *testing.T) {
	src := `<? list($a) = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListArrayIndex(t *testing.T) {
	src := `<? list($a[]) = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.ArrayDimFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListList(t *testing.T) {
	src := `<? list(list($a)) = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListEmptyItem(t *testing.T) {
	src := `<? list(, $a) = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							nil,
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestListEmptyItems(t *testing.T) {
	src := `<? list(, , $a, ) = $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							nil,
							nil,
							&expr.ArrayItem{
								Val: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
							nil,
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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
