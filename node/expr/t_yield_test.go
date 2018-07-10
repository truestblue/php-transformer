package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-transformer/node/expr"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/scalar"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestYield(t *testing.T) {
	src := `<? yield;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Yield{},
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

func TestYieldVal(t *testing.T) {
	src := `<? yield $a;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Yield{
					Value: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestYieldKeyVal(t *testing.T) {
	src := `<? yield $a => $b;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Yield{
					Key:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Value: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestYieldExpr(t *testing.T) {
	src := `<? yield 1;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Yield{
					Value: &scalar.Lnumber{Value: "1"},
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

func TestYieldKeyExpr(t *testing.T) {
	src := `<? yield $a => 1;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Yield{
					Key:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Value: &scalar.Lnumber{Value: "1"},
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

func TestYieldFrom(t *testing.T) {
	src := `<? yield from $a;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.YieldFrom{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
