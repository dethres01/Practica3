package parser

func (p *Parser) term() (*ASTNode, error) {
	// term grammar:
	// term -> <factor> <rest_of_term>
	// first token

	// New Node
	node := &ASTNode{}
	node.TokenType = "term"

	f_node, err := p.factor()
	if err != nil {
		return node, err
	}
	// add factor to AST
	node.Children = append(node.Children, f_node)
	r_node, err := p.rest_of_term()
	if err != nil {
		return node, err
	}
	// add rest_of_term to AST
	node.Children = append(node.Children, r_node)

	// append info to AST
	return node, nil
}
