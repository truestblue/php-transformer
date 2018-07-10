<<<<<<< HEAD
package scalar
=======
package scalar_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"bytes"
	"testing"

<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestDoubleQuotedScalarString(t *testing.T) {
	src := `<? "test";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"test\""},
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
func TestDoubleQuotedScalarStringWithEscapedVar(t *testing.T) {
	src := `<? "\$test";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\\$test\""},
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

func TestMultilineDoubleQuotedScalarString(t *testing.T) {
	src := `<? "
	test
	";`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "\"\n\ttest\n\t\""},
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

func TestSingleQuotedScalarString(t *testing.T) {
	src := `<? '$test';`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "'$test'"},
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

func TestMultilineSingleQuotedScalarString(t *testing.T) {
	src := `<? '
	$test
	';`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.String{Value: "'\n\t$test\n\t'"},
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
