package semantics

import (
	"Practica3/parser"
	"fmt"
)

// Semantic Analyzer function to analyze factor node
func (s *SemanticAnalyzer) Factor(node *parser.ASTNode) (string, error) {

	// factor can only have one child either a expression or a number
	if len(node.Children) == 0 {
		// return error
		return "", fmt.Errorf("factor node has no children")
	}
	// with this we can check if the factor has a expression or a number
	if node.Children[0].TokenType == "expr" {
		// expression we really just recurse into the expression until we find the number
		evalue, err := s.Expr(node.Children[0])
		if err != nil {
			return "", err
		}
		return evalue, nil
	}
	// number
	if node.Children[0].TokenType == "number" {
		// we can just add the number to the AST
		node.TokenValue = node.Children[0].TokenValue
	}
	return node.TokenValue, nil
}
