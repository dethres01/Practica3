package parser

import "fmt"

func (p *Parser) factor() error {
	//fmt.Println("Entering Factor")
	// factor grammar:
	// factor -> ( <expr> ) | <number>
	switch p.Lookahead_token {
	case "left parenthesis":
		// match left parenthesis
		err := p.Match("left parenthesis")
		if err != nil {
			return err
		}
		// match expression
		err = p.expr()
		fmt.Println("Exited Expr")
		if err != nil {
			return err
		}
		// match right parenthesis
		err = p.Match("right parenthesis")
		if err != nil {
			return err
		}
	case "number":
		// match number
		// add number to token list
		p.Tokens = append(p.Tokens, Token{
			Type: p.Lookahead_token,
			Info: p.Lookahead_value})
		err := p.Match("number")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("error on factor module")
	}
	return nil
}
