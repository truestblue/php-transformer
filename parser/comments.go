package parser

import (
<<<<<<< HEAD
	"github.com/truestblue/php-transformer/comment"
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/scanner"
=======
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/scanner"
>>>>>>> 6d554c0468596ce633490e01f7d7cb179c7dabab
)

// Comments a collection of comment groups assigned to nodes
type Comments map[node.Node][]*comment.Comment

// AddComments add comment group to the collection
func (c Comments) AddComments(node node.Node, comments []*comment.Comment) {
	c[node] = append(c[node], comments...)
}

func (c Comments) AddFromToken(node node.Node, token *scanner.Token, tokenName comment.TokenName) {
	comments := token.Comments

	for _, cmt := range comments {
		cmt.SetTokenName(tokenName)
	}

	c.AddComments(node, comments)
}

func (c Comments) AddFromChildNode(n node.Node, ch node.Node) {
	c.AddComments(n, c[ch])
	delete(c, ch)
}
