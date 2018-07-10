package scanner

import (
	"github.com/truestblue/php-transformer/comment"
	"github.com/truestblue/php-transformer/position"
)

// Token value returned by lexer
type Token struct {
	Value    string
	Position *position.Position
	Comments []*comment.Comment
}

func (t *Token) String() string {
	return string(t.Value)
}
