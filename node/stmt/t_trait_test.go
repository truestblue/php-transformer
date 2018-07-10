package stmt

import (
	"bytes"
	"testing"

	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/stmt"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
)

func TestTrait(t *testing.T) {
	src := `<? trait Foo {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Trait{
				PhpDocComment: "",
				TraitName:     &node.Identifier{Value: "Foo"},
				Stmts:         []node.Node{},
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
