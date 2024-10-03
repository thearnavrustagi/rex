package syntax

type ASTNodeType int
type Flags int

// This represents the position information about one AST node item
type Span struct {
	Start int
	End   int
}

// a comment from a regular expression with an associated span
type Comment struct {
	Span Span
	data string
}

type ASTNode interface {
	GetASTNodeType() ASTNodeType
	AddChild(ASTNode)
}

type WithComments struct {
	// the actual AST
	Root     ASTNode
	Comments Comment
}

const (
	// empty regex matches everything
	EmptyType ASTNodeType = iota + 1
	// a set of flags
	FlagType
	// a single character literal, including character sequences
	LiteralType
	// the any character character
	DotType
	// zero-width assertion
	AssertionType
	// a unicode character class
	UnicodeCharClassType
	// a perl character class
	PerlCharClassType
	// bracketed character class [a-zA-Z\pl]
	BracketCharClassType
	// repetition operator
	RepetitionType
	// grouped regular expressions
	GroupType
	// an alternation of regular expressions
	AlternationType
	// a concatenation of regular expressions
	ConcatenationType
)

type EmptyNode struct {
	span  Span
	child ASTNode
}
type DotNode struct {
	span  Span
	child ASTNode
}

type LiteralNode struct {
	span    Span
	child   ASTNode
	literal rune
}

func (EmptyNode) GetASTNodeType() ASTNodeType {
	return EmptyType
}

func (node EmptyNode) AddChild(child ASTNode) {
	node.child = child
}

func (DotNode) GetASTNodeType() ASTNodeType {
	return DotType
}

func (node DotNode) AddChild(child ASTNode) {
	node.child = child
}

func (LiteralNode) GetASTNodeType() ASTNodeType {
	return LiteralType
}

func (node LiteralNode) AddChild(child ASTNode) {
	node.child = child
}

func makeEmptyNode() *EmptyNode {
	node := new(EmptyNode)
	*node = EmptyNode{
		span: Span{
			Start: 0,
			End:   0,
		},
		child: nil,
	}
	return node
}

func makeDotNode(offset int) *DotNode {
	node := new(DotNode)
	*node = DotNode{
		span: Span{
			Start: offset,
			End:   offset + 1,
		},
		child: nil,
	}

	return node
}

func makeLiteralNode(value rune, offset int) *LiteralNode {
	node := new(LiteralNode)
	*node = LiteralNode{
		span: Span{
			Start: offset,
			End:   offset + 1,
		},
		child: nil,

		literal: value,
	}

	return node
}
