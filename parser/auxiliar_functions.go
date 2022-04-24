package parser

import (
	"errors"
	"fmt"
)

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

// separate parenthesis from operators for the moment

func ExpressionCatcher() []Token {
	// get input from user
	fmt.Println("Enter an expression: ")
	// get rune from user
	var input string
	fmt.Scanln(&input)
	// tokenize input
	tokens := Tokenize(input)
	return tokens

}
func Tokenize(input string) []Token {
	var tokens []Token

	for _, ch := range input {
		if isNumber(ch) {
			tokens = append(tokens, Token{Type: "number", Info: string(ch)})
		} else if isOperator(ch) {
			tokens = append(tokens, Token{Type: "operator", Info: string(ch)})
		} else if ch == '(' {
			tokens = append(tokens, Token{Type: "left parenthesis", Info: string(ch)})
		} else if ch == ')' {
			tokens = append(tokens, Token{Type: "right parenthesis", Info: string(ch)})
		} else {
			fmt.Println("Invalid character")
		}
	}
	tokens = append(tokens, Token{Type: "epsilon", Info: ""})
	return tokens
}
func (p *Parser) Match(terminal string) error {
	//fmt.Println(p.lookahead_token)
	if p.Lookahead_token == terminal {
		//fmt.Println("enter")
		fmt.Println(p.Lookahead_value)
		p.Counter++
		p.Lookahead_token = p.Expression[p.Counter].Type
		p.Lookahead_value = p.Expression[p.Counter].Info
		return nil
	}
	return errors.New("syntax error")
}
