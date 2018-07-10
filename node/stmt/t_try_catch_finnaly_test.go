<<<<<<< HEAD
package stmt

import (
	"bytes"
	"github.com/truestblue/php-transformer/node/expr"
	"github.com/truestblue/php-transformer/node/name"
	"testing"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
=======
package stmt_test

import (
	"bytes"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

func TestTry(t *testing.T) {
	src := `<? 
		try {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts:   []node.Node{},
				Catches: []node.Node{},
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

func TestTryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
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

func TestPhp7TryCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception|RuntimeException $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestTryCatchCatch(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
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

func TestTryCatchFinally(t *testing.T) {
	src := `<? 
		try {} catch (Exception $e) {} finally {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
				},
				Finally: &stmt.Finally{
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

func TestTryCatchCatchCatch(t *testing.T) {
	src := `<? try {} catch (Exception $e) {} catch (\RuntimeException $e) {} catch (namespace\AdditionException $e) {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				Stmts: []node.Node{},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Relative{
								Parts: []node.Node{
									&name.NamePart{Value: "AdditionException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						Stmts: []node.Node{},
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
