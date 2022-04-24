package parser

import "fmt"

func (p *Parser) term() error {
	fmt.Println("Entering Term")
	// term grammar:
	// term -> <factor> <rest_of_term>
	// first token

	err := p.factor()
	fmt.Println("Exited Factor")
	if err != nil {
		return err
	}
	err = p.rest_of_term()
	fmt.Println("Exited Rest of Term")
	if err != nil {
		return err
	}
	return nil
}
