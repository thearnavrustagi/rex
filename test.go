package main

import (
	"fmt"
	"npx/rex/syntax"
)

func main() {
	head := syntax.CompileToAST("abc.d")
	syntax.match(head, "abcxd")
	fmt.Println("hello")
}
