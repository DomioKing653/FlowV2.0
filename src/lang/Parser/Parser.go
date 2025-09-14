package Parser

import (
	"strconv"
	"strings"

	"Flow2.0/lang/Lexer"
)

/*
Parser
*/

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
	Statement
*/

func (p *Parser) Statement() Node {
	var current Lexer.Token = p.current
	if current.Type == Lexer.LET {
		p.nextToken()
		if p.current.Type != Lexer.IDENTIFIER {
			panic("Expected IDENTIFIER but found" + current.Type)
		}
		ident := p.current.Value
		p.nextToken()
		if current.Type == Lexer.EQUALS {
			panic("Expected EQUALS but found" + current.Type)
		}
		p.nextToken()
		var expression Node = p.expr()
		return VariableNode{Name: ident, Value: expression}
	}
	if current.Type == Lexer.IDENTIFIER {
		ident := current.Value
		p.nextToken()
		if p.current.Type == Lexer.EQUALS {
			p.nextToken()
			exprNode := p.expr()
			return VariableAssignNode{Name: ident, Value: exprNode}
		}
		return VariableAccessNode{Name: ident}
	}
	return p.expr()
}

/*
Factor
*/
func (p *Parser) factor() Node {
	tok := p.current
	if tok.Type == Lexer.INT {
		val, _ := strconv.ParseFloat(tok.Value, 64)
		p.nextToken()
		return NumberNode{Value: val}
	}
	if tok.Type == Lexer.FLOAT {
		var valueStr string = strings.ReplaceAll(tok.Value, ",", ".")
		val, _ := strconv.ParseFloat(valueStr, 64)
		p.nextToken()
		return NumberNode{Value: val}
	} else {
		panic("Syntax Error: Expected INT  or FLOAT but found " + tok.Type)
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
	program := ProgramNode{}
	for {
		if p.current.Type == Lexer.EOF {
			return program
		}
		program.statements = append(program.statements, p.Statement())
	}
}
