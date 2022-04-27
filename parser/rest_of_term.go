package parser

import "fmt"

func (p *Parser) rest_of_term() (*ASTNode, error) {
	// rest_of_term grammar:
	// rest_of_term -> * <factor> <rest_of_term> | / <factor> <rest_of_term> | epsilon
	// first token

	// New Node
	node := &ASTNode{}
	node.TokenType = "rest_of_term"

	switch p.Lookahead_token {
	case "operator":
		// match operator
		// verify symbol
		// new token
		var aux Token
		aux.Type = p.Lookahead_token
		aux.Info = p.Lookahead_value
		op_node := &ASTNode{
			TokenType:  "operator",
			TokenValue: p.Lookahead_value,
		}
		if p.Lookahead_value != "*" && p.Lookahead_value != "/" {
			// this means we entered + or - which is part of rest_of_expr
			// meaning we have to exit this function
			//fmt.Println("Got operator: ", p.Lookahead_value)
			return node, nil
		}
		// add operator to AST
		node.Children = append(node.Children, op_node)
		err := p.Match("operator")
		if err != nil {
			return node, err
		}
		// add operator to AST
		// match factor
		f_node, err := p.factor()
		fmt.Println("Exited Factor")
		if err != nil {
			return node, err
		}
		// add factor to AST
		node.Children = append(node.Children, f_node)
		// match rest_of_term
		p.Tokens = append(p.Tokens, aux)
		r_node, err := p.rest_of_term()
		//fmt.Println("Exited Rest of Term")
		if err != nil {
			return node, err
		}
		// add rest_of_term to AST
		node.Children = append(node.Children, r_node)

	case "epsilon":
		// exit
	case "right parenthesis":

	default:
		return node, fmt.Errorf("error on rest_of_term module")
	}
	return node, nil
}
