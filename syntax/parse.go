package syntax

func CompileToAST(expr string) ASTNode {
	runes := []rune(expr)
	itr := 0
	head := makeEmptyNode()
	var ast_itr ASTNode = head

	for itr < len(runes) {
		delta, child := getCorrespondingNode(runes[itr:], itr)
		itr = itr + delta
		ast_itr.AddChild(child)
		ast_itr = child
	}

	return head.child
}

func getCorrespondingNode(expr []rune, offset int) (delta int, node ASTNode) {
	switch char := expr[0]; char {
	case DOT_CHARACTER:
		delta = 1
		node = makeDotNode(offset)
	default:
		delta = 1
		node = makeLiteralNode(char, offset)
	}
	return
}
