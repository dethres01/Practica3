package main

// grammar:
/*
expr -> <term> <rest_of_expr>
rest_of_expr -> + <term> <rest_of_expr> | - <term> <rest_of_expr> | epsilon
term -> <factor> <rest_of_term>
rest_of_term -> * <factor> <rest_of_term> | / <factor> <rest_of_term> | epsilon
factor -> ( <expr> ) | <number>
<number> -> 1 | 2 | 3 | ...
*/
import (
	"Practica3/calculator"
	"Practica3/parser"
	"fmt"
)

// Token Object

func main() {
	p := parser.NewParser()
	p.Expression = parser.ExpressionCatcher()
	err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Expression)
	fmt.Println(p.Tokens)

	// resolve the expression
	res, err := calculator.Resolve(p.Tokens)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
