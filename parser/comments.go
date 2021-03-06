package parser

import (
	"github.com/truestblue/php-transformer/comment"
	"github.com/truestblue/php-transformer/node"
	"github.com/truestblue/php-transformer/scanner"
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
