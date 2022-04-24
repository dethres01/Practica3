package parser

import "fmt"

func (p *Parser) expr() error {
	fmt.Println("Entering Expr")

	// expr grammar:
	// expr -> <term> <rest_of_expr>
	// first token
	err := p.term()
	fmt.Println("Exited Term")
	if err != nil {
		return err
	}
	err = p.rest_of_expr()
	fmt.Println("Exited Rest of Expr")
	if err != nil {
		return err
	}
	return nil
}
