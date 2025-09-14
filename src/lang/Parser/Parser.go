package Parser

import (
	"strconv"

	"Flow2.0/lang/Lexer"
)

type Parser struct {
	current Lexer.Token
	pos     int
	tokens  []Lexer.Token
}

func NewParser(tokens []Lexer.Token) *Parser {
	p := &Parser{tokens: tokens, pos: -1}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.pos++
	if p.pos < len(p.tokens) {
		p.current = p.tokens[p.pos]
	} else {
		p.current = Lexer.Token{Type: Lexer.EOF, Value: ""}
	}
}

/*
* Factor
 */
func (p *Parser) factor() Node {
	tok := p.current

	if tok.Type == Lexer.NUMBER {
		val, _ := strconv.Atoi(tok.Value)
		p.nextToken()
		return NumberNode{Value: val}
	} else {
		panic("Syntax Error: Expected INT but found " + tok.Type)
	}
}

/*
* Term
 */
func (p *Parser) term() Node {
	node := p.factor()
	for p.current.Type == Lexer.STAR || p.current.Type == Lexer.SLASH {
		op := p.current
		p.nextToken()
		node = BinaryOperationNode{
			Left:     node,
			Operator: op.Value,
			Right:    p.factor(),
		}
	}
	return node
}

/*
* Expression
 */
func (p *Parser) expr() Node {
	node := p.term()
	for p.current.Type == Lexer.PLUS || p.current.Type == Lexer.MINUS {
		op := p.current
		p.nextToken()
		node = BinaryOperationNode{
			Left:     node,
			Operator: op.Value,
			Right:    p.term(),
		}
	}
	return node
}
func (p *Parser) Parse() Node {
	return p.expr()
}
