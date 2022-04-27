package semantics

import (
	"Practica3/parser"
	"fmt"
)

// semantic analyzer function to analyze rest_of_term node
func (s *SemanticAnalyzer) Rest_of_term(node *parser.ASTNode) (string, error) {
	// grammar:
	// rest_of_term -> * <factor> <rest_of_term> | / <factor> <rest_of_term> | e
	// children
	// epsilon check
	if len(node.Children) == 0 {
		// return epsilon
		return "", nil
	}
	// skip if operator is not a * or /
	if node.Children[0].TokenValue != "*" && node.Children[0].TokenValue != "/" {
		return "", nil
	}
	// add operator to AST
	operator := node.Children[0].TokenValue
	// go to factor
	fvalue, err := s.Factor(node.Children[1])
	if err != nil {
		return "", err
	}

	// Evaluate the rest of term
	rvalue, err := s.Rest_of_term(node.Children[2])
	if err != nil {
		return "", err
	}
	node.TokenValue = fmt.Sprintf("%s %s %s", operator, fvalue, rvalue)

	return node.TokenValue, nil
}
