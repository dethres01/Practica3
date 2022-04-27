package parser

import "fmt"

func (p *Parser) factor() (*ASTNode, error) {
	//fmt.Println("Entering Factor")
	// factor grammar:
	// factor -> ( <expr> ) | <number>

	// New Node
	node := &ASTNode{}
	node.TokenType = "factor"

	switch p.Lookahead_token {
	case "left parenthesis":
		// match left parenthesis
		err := p.Match("left parenthesis")
		if err != nil {
			return node, err
		}
		// match expression
		e_node, err := p.expr()
		fmt.Println("Exited Expr")
		if err != nil {
			return node, err
		}
		// add expression to AST
		node.Children = append(node.Children, e_node)
		err = p.Match("right parenthesis")
		if err != nil {
			return node, err
		}
	case "number":
		// match number
		// add number to token list
		p.Tokens = append(p.Tokens, Token{
			Type: p.Lookahead_token,
			Info: p.Lookahead_value})
		// Add Child
		node.Children = append(node.Children, &ASTNode{
			TokenType:  p.Lookahead_token,
			TokenValue: p.Lookahead_value,
			Children: []*ASTNode{
				&ASTNode{
					TokenType:  p.Lookahead_token,
					TokenValue: p.Lookahead_value,
				},
			},
		})

		err := p.Match("number")
		if err != nil {
			return node, err
		}
	default:
		return node, fmt.Errorf("error on factor module")
	}
	return node, nil
}
