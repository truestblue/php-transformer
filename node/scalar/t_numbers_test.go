package scalar_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		diff := pretty.Compare(expected, actual)

		if diff != "" {
			t.Errorf("diff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("expected and actual are not equal\n")
		}

	}
}

func TestLNumber(t *testing.T) {
	src := `<? 1234567890123456789;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "1234567890123456789"},
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

func TestDNumber(t *testing.T) {
	src := `<? 12345678901234567890;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "12345678901234567890"},
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

func TestFloat(t *testing.T) {
	src := `<? 0.;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0."},
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

func TestBinaryLNumber(t *testing.T) {
	src := `<? 0b0111111111111111111111111111111111111111111111111111111111111111;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0b0111111111111111111111111111111111111111111111111111111111111111"},
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

func TestBinaryDNumber(t *testing.T) {
	src := `<? 0b1111111111111111111111111111111111111111111111111111111111111111;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0b1111111111111111111111111111111111111111111111111111111111111111"},
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

func TestHLNumber(t *testing.T) {
	src := `<? 0x007111111111111111;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Lnumber{Value: "0x007111111111111111"},
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

func TestHDNumber(t *testing.T) {
	src := `<? 0x8111111111111111;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &scalar.Dnumber{Value: "0x8111111111111111"},
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
