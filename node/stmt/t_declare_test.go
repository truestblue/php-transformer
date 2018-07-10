<<<<<<< HEAD
package stmt
=======
package stmt_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"bytes"
	"testing"

<<<<<<< HEAD
	"github.com/z7zmey/php-transformer/node/scalar"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
=======
	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestDeclare(t *testing.T) {
	src := `<? declare(ticks=1);`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.Nop{},
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

func TestDeclareStmts(t *testing.T) {
	src := `<? declare(ticks=1, strict_types=1) {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "strict_types"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.StmtList{
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

func TestAltDeclare(t *testing.T) {
	src := `<? declare(ticks=1): enddeclare;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Declare{
				Consts: []node.Node{
					&stmt.Constant{
						PhpDocComment: "",
						ConstantName:  &node.Identifier{Value: "ticks"},
						Expr:          &scalar.Lnumber{Value: "1"},
					},
				},
				Stmt: &stmt.StmtList{
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
