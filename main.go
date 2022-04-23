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
import "Practica3/parser"

// Token Object

func main() {
	p := parser.NewParser()
	p.Expression = parser.ExpressionCatcher()
}
