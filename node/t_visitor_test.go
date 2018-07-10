package node_test

import (
	"reflect"
	"testing"

	"github.com/truestblue/php-transformer/node/stmt"

	"github.com/truestblue/php-transformer/node/scalar"
	"github.com/truestblue/php-transformer/walker"

	"github.com/truestblue/php-transformer/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/truestblue/php-transformer/node"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&node.Identifier{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&node.Nullable{Expr: &expr.Variable{VarName: &node.Identifier{Value: "foo"}}},
		[]string{"Expr"},
		nil,
	},
	{
		&node.Argument{Variadic: true, Expr: &expr.Variable{VarName: &node.Identifier{Value: "foo"}}},
		[]string{"Expr"},
		map[string]interface{}{"IsReference": false, "Variadic": true},
	},
	{
		&node.Parameter{
			ByRef:        false,
			Variadic:     true,
			VariableType: &node.Identifier{Value: "foo"},
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "bar"}},
			DefaultValue: &scalar.Lnumber{Value: "0"},
		},
		[]string{"VariableType", "Variable", "DefaultValue"},
		map[string]interface{}{"ByRef": false, "Variadic": true},
	},
	{
		&node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{},
			},
		},
		[]string{"Arguments"},
		map[string]interface{}{},
	},
	{
		&node.Root{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		map[string]interface{}{},
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
	for _, tt := range nodesToTest {
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
	for _, tt := range nodesToTest {
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
	for _, tt := range nodesToTest {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}
