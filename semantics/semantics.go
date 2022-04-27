package semantics

import (
	"Practica3/parser"
	"fmt"
)

type SemanticAnalyzer struct {
	AST *parser.AST
}

func NewSemanticAnalyzer(p *parser.AST) *SemanticAnalyzer {
	return &SemanticAnalyzer{
		AST: p,
	}
}

func (s *SemanticAnalyzer) Analyze() error {
	// we basically want to resolve the AST so we can resolve the operation of the expression
	// we can do some basic tree fixing here
	// first child is the root node

	s.AST.Root = s.AST.Children[0]

	// Right now Root Info doesn't have the result of the expression so we need to get into the children

	// we need to check if the root has children
	if len(s.AST.Root.Children) == 0 {
		// return error
		return fmt.Errorf("root node has no children")
	}
	tvalue, err := s.Term(s.AST.Root.Children[0])
	if err != nil {
		return err
	}
	fmt.Println("Result: ", tvalue)

	// rest of expr
	rvalue, err := s.Rest_of_expr(s.AST.Root.Children[1])
	if err != nil {
		return err
	}
	fmt.Println("Result: ", rvalue)

	// real result
	s.AST.Root.TokenValue = fmt.Sprintf("%s %s", tvalue, rvalue)
	fmt.Println("Result: ", s.AST.Root.TokenValue)
	s.AST.Root.TokenValue = s.parseString(s.AST.Root.TokenValue)
	fmt.Println("Result: ", s.AST.Root.TokenValue)

	return nil

}
