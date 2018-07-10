package errors

import (
	"reflect"
	"testing"

	"github.com/truestblue/php-transformer/position"

	"github.com/truestblue/php-transformer/errors"
	"github.com/truestblue/php-transformer/scanner"

	"github.com/kylelemons/godebug/pretty"
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

func TestConstructor(t *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)
	token := &scanner.Token{
		Value:    `test`,
		Position: pos,
	}

	actual := errors.NewError("message", token)

	expected := &errors.Error{
		Msg: "message",
		Pos: pos,
	}

	assertEqual(t, expected, actual)
}

func TestPrint(t *testing.T) {
	pos := position.NewPosition(1, 2, 3, 4)
	token := &scanner.Token{
		Value:    `test`,
		Position: pos,
	}

	Error := errors.NewError("message", token)

	actual := Error.String()

	expected := "message at line 1"

	assertEqual(t, expected, actual)
}
