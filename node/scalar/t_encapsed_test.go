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

func TestSimpleVar(t *testing.T) {
	src := `<? "test $var";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "var"}},
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

func TestSimpleVarOneChar(t *testing.T) {
	src := `<? "test $a";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestSimpleVarEndsEcapsed(t *testing.T) {
	src := `<? "test $var\"";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "var"}},
						&scalar.EncapsedStringPart{Value: "\\\""},
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

func TestStringVarCurveOpen(t *testing.T) {
	src := `<? "=$a{$b}";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "="},
						&expr.Variable{VarName: &node.Identifier{Value: "a"}},
						&expr.Variable{VarName: &node.Identifier{Value: "b"}},
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

func TestSimpleVarPropertyFetch(t *testing.T) {
	src := `<? "test $foo->bar()";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.PropertyFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Property: &node.Identifier{Value: "bar"},
						},
						&scalar.EncapsedStringPart{Value: "()"},
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

func TestDollarOpenCurlyBraces(t *testing.T) {
	src := `<? "test ${foo}";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.Variable{VarName: &node.Identifier{Value: "foo"}},
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

func TestDollarOpenCurlyBracesDimNumber(t *testing.T) {
	src := `<? "test ${foo[0]}";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.ArrayDimFetch{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Dim:      &scalar.Lnumber{Value: "0"},
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

func TestCurlyOpenMethodCall(t *testing.T) {
	src := `<? "test {$foo->bar()}";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Encapsed{
					Parts: []node.Node{
						&scalar.EncapsedStringPart{Value: "test "},
						&expr.MethodCall{
							Variable:     &expr.Variable{VarName: &node.Identifier{Value: "foo"}},
							Method:       &node.Identifier{Value: "bar"},
							ArgumentList: &node.ArgumentList{},
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
