package semantics

import (
	"Practica3/parser"
	"fmt"
)

// SEMANTIC ANALYZER FUNCTION TO ANALYZE TERM NODE
func (s *SemanticAnalyzer) Term(node *parser.ASTNode) (string, error) {

	// factor
	fvalue, err := s.Factor(node.Children[0])
	if err != nil {
		return "", err
	}
	// add value to AST
	node.TokenValue = fvalue

	// rest of term
	rvalue, err := s.Rest_of_term(node.Children[1])
	if err != nil {
		return "", err
	}
	node.TokenValue = fmt.Sprintf("%s %s", node.TokenValue, rvalue)

	// should be able to resolve this term
	//node.TokenValue = s.parseString(node.TokenValue)

	return node.TokenValue, nil
}
