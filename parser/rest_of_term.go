package parser

import "fmt"

func (p *Parser) rest_of_term() error {
	fmt.Println("Entering Rest of Term")
	// rest_of_term grammar:
	// rest_of_term -> * <factor> <rest_of_term> | / <factor> <rest_of_term> | epsilon
	// first token
	fmt.Println("Lookahead: ", p.Lookahead_token)
	switch p.Lookahead_token {
	case "operator":
		// match operator
		// verify symbol
		// new token
		var aux Token
		aux.Type = p.Lookahead_token
		aux.Info = p.Lookahead_value
		if p.Lookahead_value != "*" && p.Lookahead_value != "/" {
			// this means we entered + or - which is part of rest_of_expr
			// meaning we have to exit this function
			fmt.Println("Got operator: ", p.Lookahead_value)
			return nil
		}
		err := p.Match("operator")
		if err != nil {
			return err
		}
		// match factor
		err = p.factor()
		fmt.Println("Exited Factor")
		if err != nil {
			return err
		}
		// match rest_of_term
		p.Tokens = append(p.Tokens, aux)
		err = p.rest_of_term()
		//fmt.Println("Exited Rest of Term")
		if err != nil {
			return err
		}

	case "epsilon":
		// exit
	case "right parenthesis":

	default:
		return fmt.Errorf("error on rest_of_term module")
	}
	return nil
}
