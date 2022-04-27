package semantics

import (
	"Practica3/parser"
	"fmt"
)

// semantic analyzer function to analyze expr node
func (s *SemanticAnalyzer) Expr(node *parser.ASTNode) (string, error) {

	// evaluate the term
	tvalue, err := s.Term(node.Children[0])
	if err != nil {
		return "", err
	}
	// add value to AST
	node.TokenValue = tvalue
	rvalue, err := s.Rest_of_expr(node.Children[1])
	if err != nil {
		return "", err
	}

	node.TokenValue = fmt.Sprintf("%s %s", node.TokenValue, rvalue)
	// should be able to resolve this expr
	node.TokenValue = s.parseString(node.TokenValue)
	return node.TokenValue, nil
}
