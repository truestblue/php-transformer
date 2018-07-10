<<<<<<< HEAD
package scalar
=======
package scalar_test
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/walker"
=======
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

var nameNodesTests = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&scalar.String{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&scalar.Lnumber{Value: "0"},
		[]string{},
		map[string]interface{}{"Value": "0"},
	},
	{
		&scalar.Dnumber{Value: "0.1"},
		[]string{},
		map[string]interface{}{"Value": "0.1"},
	},
	{
		&scalar.MagicConstant{Value: "__DIR__"},
		[]string{},
		map[string]interface{}{"Value": "__DIR__"},
	},
	{
		&scalar.EncapsedStringPart{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&scalar.Encapsed{Parts: []node.Node{&scalar.EncapsedStringPart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&scalar.Heredoc{Label: "LBL", Parts: []node.Node{&scalar.EncapsedStringPart{Value: "foo"}}},
		[]string{"Parts"},
		map[string]interface{}{"Label": "LBL"},
	},
}

type visitorMock struct {
	visitChildren bool
	visitedKeys   []string
}

func (v *visitorMock) EnterNode(n walker.Walkable) bool { return v.visitChildren }
func (v *visitorMock) GetChildrenVisitor(key string) walker.Visitor {
	v.visitedKeys = append(v.visitedKeys, key)
	return &visitorMock{v.visitChildren, nil}
}
func (v *visitorMock) LeaveNode(n walker.Walkable) {}

func TestNameVisitorDisableChildren(t *testing.T) {
	for _, tt := range nameNodesTests {
		v := &visitorMock{false, nil}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

func TestNameVisitor(t *testing.T) {
	for _, tt := range nameNodesTests {
		v := &visitorMock{true, nil}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nameNodesTests {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}
