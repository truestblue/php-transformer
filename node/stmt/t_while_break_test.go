package stmt_test

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node/scalar"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestBreakEmpty(t *testing.T) {
	src := `<? while (1) { break; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{nil},
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

func TestBreakLight(t *testing.T) {
	src := `<? while (1) { break 2; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{
							Expr: &scalar.Lnumber{Value: "2"},
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

func TestBreak(t *testing.T) {
	src := `<? while (1) : break(3); endwhile;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.AltWhile{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.StmtList{
					Stmts: []node.Node{
						&stmt.Break{
							Expr: &scalar.Lnumber{Value: "3"},
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
