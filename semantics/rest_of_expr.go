package semantics

import (
	"Practica3/parser"
	"fmt"
)

// semantic analyzer function to analyze rest_of_expr node
func (s *SemanticAnalyzer) Rest_of_expr(node *parser.ASTNode) (string, error) {
	// grammar:
	// rest_of_expr -> + <term> <rest_of_expr> | - <term> <rest_of_expr> | e

	// children
	// epsilon check
	if len(node.Children) == 0 {
		// return epsilon
		return "", nil
	}
	// add operator to AST
	operator := node.Children[0].TokenValue
	// go to term
	tvalue, err := s.Term(node.Children[1])
	if err != nil {
		return "", err
	}

	// Evaluate the rest of expr
	rvalue, err := s.Rest_of_expr(node.Children[2])
	if err != nil {
		return "", err
	}
	node.TokenValue = fmt.Sprintf("%s %s %s", operator, tvalue, rvalue)

	return node.TokenValue, nil
}
