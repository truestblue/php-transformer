/*

A Parser for PHP written in Go

Features:

	* Fully support PHP5 and PHP7 syntax
	* Abstract syntax tree representation
	* Traversing AST
	* Namespace resolver

Install:

	go get github.com/truestblue/php-transformer

CLI dumper:

	$GOPATH/bin/php-transformer -php5 /path/to/file/or/dir


Package usage example:

	package main

	import (
		"fmt"
		"bytes"
		"os"

		"github.com/truestblue/php-transformer/php7"
		"github.com/truestblue/php-transformer/visitor"

	)

	func main() {
		src := bytes.NewBufferString(`<? echo "Hello world";`)

		parser := php7.NewParser(src, "example.php")
		parser.Parse()

		for _, e := range parser.GetErrors() {
			fmt.Println(e)
		}

		visitor := visitor.Dumper{
			Writer:    os.Stdout,
			Indent:    "",
			Comments:  parser.GetComments(),
			Positions: parser.GetPositions(),
		}

		rootNode := parser.GetRootNode()
		rootNode.Walk(visitor)
	}
*/
package main // import "github.com/truestblue/php-transformer"
