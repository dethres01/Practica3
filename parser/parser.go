package parser

type Token struct {
	Type string
	Info string
}
type Parser struct {
	Expression      []Token
	Counter         int
	Lookahead_token string
	Lookahead_value string
	Tokens          []Token
}

func NewParser() *Parser {
	return &Parser{
		Counter: 0,
	}
}
func (p *Parser) Parse() (*AST, error) {
	// start parsing
	// initial grammar:
	// expr -> <term> <rest_of_expr>

	// first token
	p.Lookahead_token = p.Expression[p.Counter].Type
	p.Lookahead_value = p.Expression[p.Counter].Info

	// initialize AST
	ast := &AST{}
	e_node, err := p.expr()
	if err != nil {
		return ast, err
	}
	ast.Children = append(ast.Children, e_node)

	return ast, nil
}

//Since we cannot use a stack we have to transform the entire code to have a proper AST
