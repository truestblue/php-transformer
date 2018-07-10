package stmt

import (
	"bytes"
	"github.com/z7zmey/php-transformer/node/name"
	"testing"

	"github.com/z7zmey/php-transformer/node"
	"github.com/z7zmey/php-transformer/node/stmt"
	"github.com/z7zmey/php-transformer/php5"
	"github.com/z7zmey/php-transformer/php7"
)

func TestNamespace(t *testing.T) {
	src := `<? namespace Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
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

func TestNamespaceStmts(t *testing.T) {
	src := `<? namespace Foo {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				Stmts: []node.Node{},
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

func TestAnonymousNamespace(t *testing.T) {
	src := `<? namespace {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Namespace{
				Stmts: []node.Node{},
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
