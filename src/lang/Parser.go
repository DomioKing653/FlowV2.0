package lang

type Parser struct {
	_currentToken Token
	_TokenIdx     int
	_listTokens   []Token
}

func (p *Parser) NextToken() {
	p._TokenIdx++
	p._currentToken = p._listTokens[p._TokenIdx]
}
func (p *Parser) Parser() {
}
