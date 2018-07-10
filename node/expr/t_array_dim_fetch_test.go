package expr

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/truestblue/php-transformer/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/truestblue/php-transformer/node/scalar"

	"github.com/truestblue/php-transformer/node"
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

func TestArrayDimFetch(t *testing.T) {
	src := `<? $a[1];`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Dim:      &scalar.Lnumber{Value: "1"},
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

func TestArrayDimFetchNested(t *testing.T) {
	src := `<? $a[1][2];`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ArrayDimFetch{
					Variable: &expr.ArrayDimFetch{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Dim:      &scalar.Lnumber{Value: "1"},
					},
					Dim: &scalar.Lnumber{Value: "2"},
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
