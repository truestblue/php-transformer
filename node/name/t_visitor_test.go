package name_test

import (
	"reflect"
	"testing"

	"github.com/truestblue/php-transformer/node/name"
	"github.com/truestblue/php-transformer/walker"

	"github.com/kylelemons/godebug/pretty"
	"github.com/truestblue/php-transformer/node"
)

var nameNodesTests = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.FullyQualified{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.Relative{Parts: []node.Node{&name.NamePart{Value: "foo"}}},
		[]string{"Parts"},
		nil,
	},
	{
		&name.NamePart{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
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
