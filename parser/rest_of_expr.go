package parser

import "fmt"

func (p *Parser) rest_of_expr() error {
	//fmt.Println("Entering Rest of Expr")
	// rest_of_expr grammar:
	// rest_of_expr -> + <term> <rest_of_expr> | - <term> <rest_of_expr> | epsilon
	// first token
	switch p.Lookahead_token {
	case "operator":
		// verify symbol
		// new token
		var aux Token
		aux.Type = p.Lookahead_token
		aux.Info = p.Lookahead_value
		if p.Lookahead_value != "+" && p.Lookahead_value != "-" {
			return fmt.Errorf("error on rest_of_expr module symbol: %s", p.Lookahead_value)
		}
		// match operator
		err := p.Match("operator")
		if err != nil {
			return err
		}
		// match term
		err = p.term()
		//fmt.Println("Exited Term")
		if err != nil {
			return err
		}
		// match rest_of_expr
		p.Tokens = append(p.Tokens, aux)
		err = p.rest_of_expr()
		//fmt.Println("Exited Rest of Expr")
		if err != nil {
			return err
		}
	case "epsilon":
	}
	// exit
	return nil
}
