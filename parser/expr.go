package parser

func (p *Parser) expr() (*ASTNode, error) {

	// expr grammar:
	// expr -> <term> <rest_of_expr>
	// first token
	// New AST
	// New Node
	node := &ASTNode{}
	node.TokenType = "expr"
	t_node, err := p.term()
	if err != nil {
		return node, err
	}
	// add term to AST
	node.Children = append(node.Children, t_node)
	r_node, err := p.rest_of_expr()
	if err != nil {
		return node, err
	}
	// add rest_of_expr to AST
	node.Children = append(node.Children, r_node)

	// append info to AST
	return node, nil
}
