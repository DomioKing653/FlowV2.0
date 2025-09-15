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
	var current = p.current
	switch p.current.Type {
	case Lexer.CONST, Lexer.LET:
		varType := p.current.Type
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
		var expression = p.expr()
		if varType == Lexer.CONST {
			return VariableNode{Name: ident, Value: expression, Constant: true}
		}
		return VariableNode{Name: ident, Value: expression, Constant: false}

	case Lexer.IDENTIFIER:
		ident := current.Value
		p.nextToken()
		if p.current.Type == Lexer.EQUALS {
			p.nextToken()
			exprNode := p.expr()
			return VariableAssignNode{Name: ident, Value: exprNode}
		}
		return VariableAccessNode{Name: ident}
	case Lexer.PRINTLN:
		p.nextToken()
		if p.current.Type != Lexer.LPAREN {
			panic("Expected LPAREN but found" + current.Type)
		}
		p.nextToken()
		exprNode := p.expr()
		if p.current.Type != Lexer.RPAREN {
			panic("Expected RPAREN but found" + current.Type)
		}
		p.nextToken()
		return PrintLnNode{exprNode}
	default:
		return p.expr()
	}
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
		var valueStr = strings.ReplaceAll(tok.Value, ",", ".")
		val, _ := strconv.ParseFloat(valueStr, 64)
		p.nextToken()
		return NumberNode{Value: val}

	}
	if tok.Type == Lexer.IDENTIFIER {
		ident := tok.Value
		p.nextToken()
		return VariableAccessNode{Name: ident}
	}
	if tok.Type == Lexer.LPAREN {
		p.nextToken()
		exprNode := p.expr()
		if tok.Type == Lexer.RPAREN {
			panic("Expected RPAREN but found" + tok.Value)
		}
		p.nextToken()
		return exprNode
	} else {
		panic("Syntax Error: Expected VALUE but found " + tok.Type)
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
			break
		}
		program.statements = append(program.statements, p.Statement())
	}
	return program
}
