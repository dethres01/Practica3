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
func (p *Parser) Parse() error {
	// start parsing
	// initial grammar:
	// expr -> <term> <rest_of_expr>

	// first token
	p.Lookahead_token = p.Expression[p.Counter].Type
	p.Lookahead_value = p.Expression[p.Counter].Info
	err := p.expr()
	if err != nil {
		return err
	}
	return nil
}
