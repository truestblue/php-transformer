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

//func (t *Token) String() string {
//	fmt.Println(t.Value)
//	return string(t.Value)
//}
