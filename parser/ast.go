package parser

// here we have the AST and it's methods

// Can have more than 2 children
type AST struct {
	Root     *ASTNode
	Children []*ASTNode
}

type ASTNode struct {
	TokenType  string
	TokenValue string
	Children   []*ASTNode
}
