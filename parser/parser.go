package parser

import (
	"errors"
	"fmt"
)

type Token struct {
	Type string
	Info string
}
type Parser struct {
	Expression      []Token
	Counter         int
	Lookahead_token string
	Tokens          []Token
}

func NewParser() *Parser {
	return &Parser{
		Counter: 0,
	}
}

func (p *Parser) Remain() error {
	// basically here we enter a loop until we have an "epilson" token
	// in reality, every time we enter remain it should be with an operator token
	switch p.Lookahead_token {
	case "operator":
		// save operator token on auxiliar object
		aux := p.Expression[p.Counter]
		// match operator token
		err := p.Match("operator")
		if err != nil {
			return err
		}
		err = p.term()
		if err != nil {
			return err
		}
		// add operator token to list of tokens
		p.Tokens = append(p.Tokens, aux)
		// call remain again
		err = p.Remain()
		if err != nil {
			return err
		}
	case "epsilon":
		if p.Counter%2 == 0 {
			return errors.New("syntax error, expected operator, got epsilon")
		}
	default:
		return errors.New("syntax error, expected operator, got " + p.Lookahead_token)
	}
	return nil

}
func (p *Parser) Match(terminal string) error {
	//fmt.Println(p.lookahead_token)
	if p.Lookahead_token == terminal {
		//fmt.Println("enter")
		p.Counter++
		p.Lookahead_token = p.Expression[p.Counter].Type
		return nil
	}
	return errors.New("syntax error")
}
func (p *Parser) term() error {

	err := p.Match("term")
	if err != nil {
		return fmt.Errorf("syntax error, expected term, got %s", p.Lookahead_token)
	}
	p.Tokens = append(p.Tokens, p.Expression[p.Counter-1])
	return nil
}
