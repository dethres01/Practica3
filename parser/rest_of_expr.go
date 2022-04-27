package parser

import "fmt"

func (p *Parser) rest_of_expr() (*ASTNode, error) {
	//fmt.Println("Entering Rest of Expr")
	// rest_of_expr grammar:
	// rest_of_expr -> + <term> <rest_of_expr> | - <term> <rest_of_expr> | epsilon
	// first token

	// New Node
	node := &ASTNode{}
	node.TokenType = "rest_of_expr"
	switch p.Lookahead_token {
	case "operator":
		// verify symbol
		// new token
		var aux Token
		aux.Type = p.Lookahead_token
		aux.Info = p.Lookahead_value
		op_node := &ASTNode{
			TokenType:  "operator",
			TokenValue: p.Lookahead_value,
		}
		if p.Lookahead_value != "+" && p.Lookahead_value != "-" {
			return node, fmt.Errorf("error on rest_of_expr module symbol: %s", p.Lookahead_value)
		}
		// add operator to AST
		node.Children = append(node.Children, op_node)
		// match operator
		err := p.Match("operator")
		if err != nil {
			return node, err
		}
		// match term
		t_node, err := p.term()
		//fmt.Println("Exited Term")
		if err != nil {
			return node, err
		}
		// add term to AST
		node.Children = append(node.Children, t_node)
		// match rest_of_expr
		p.Tokens = append(p.Tokens, aux)
		r_node, err := p.rest_of_expr()
		//fmt.Println("Exited Rest of Expr")
		if err != nil {
			return node, err
		}
		// add rest_of_expr to AST
		node.Children = append(node.Children, r_node)
	case "epsilon":
	}
	// exit
	return node, nil
}
