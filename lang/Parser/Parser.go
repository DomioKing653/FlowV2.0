package Parser

import (
    "Flow2.0/lang/Lexer"
    "Flow2.0/lang/shared"
    "errors"
    "strconv"
    "strings"
)

/*
	Parser
*/

type Parser struct {
	current Lexer.Token
	pos     int
	tokens  []Lexer.Token
}

func NewParser(tokens []Lexer.Token) shared.IParser {
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

func (p *Parser) Statement() (shared.Node,error) {

	var current = p.current
	switch p.current.Type {
	/*
		Variable
	*/
	case Lexer.CONST, Lexer.LET:
		varType := p.current.Type
		p.nextToken()
		if p.current.Type != Lexer.IDENTIFIER {
			return nil,errors.New("Expected IDENTIFIER but found:" + p.current.Value)
		}
		ident := p.current.Value
		p.nextToken()
		if current.Type == Lexer.EQUALS {
			return nil,errors.New("Expected EQUALS but found:" + p.current.Value)
		}
		p.nextToken()
		var expression = p.expr()
		if varType == Lexer.CONST {
			return VariableNode{Name: ident, Value: expression, Constant: true},nil
		}
		return VariableNode{Name: ident, Value: expression, Constant: false},nil
	/*
		Identifier
	*/
	case Lexer.IDENTIFIER:
		ident := current.Value
		p.nextToken()
		if p.current.Type == Lexer.EQUALS {
			p.nextToken()
			exprNode := p.expr()
			return VariableAssignNode{Name: ident, Value: exprNode},nil
		}
		if p.current.Type == Lexer.LPAREN {
			p.nextToken()
		}
		return VariableAccessNode{Name: ident},nil
	/*
		Println
	*/
	case Lexer.PRINTLN:
		p.nextToken()
		if p.current.Type != Lexer.LPAREN {
			return nil,errors.New("Expected LPAREN but found:" + p.current.Value)
		}
		p.nextToken()
		exprNode := p.expr()
		if p.current.Type != Lexer.RPAREN {
			return nil,errors.New("Expected RPAREN but found:" + p.current.Value)
		}
		p.nextToken()
		return PrintLnNode{exprNode},nil
	/*
		Loop
	*/
	case Lexer.LOOP:
		p.nextToken()
		var listNode []shared.Node
		if p.current.Type != Lexer.OpeningParen {
			return nil,errors.New("Expected OPPENING PAREN but found:" + p.current.Value)
		}
		p.nextToken()
		for {
			if p.current.Type == Lexer.ClosingParen {
				break
			} else {
				stmt,err:=p.Statement()
				shared.Check(err)
				listNode = append(listNode, stmt)
			}
		}
		p.nextToken()
		return LoopNode{Nodes: listNode},nil
	/*
		IF
	*/
	case Lexer.IF:
		p.nextToken()
		if p.current.Type != Lexer.LPAREN{
			return nil,errors.New("Expected LPAREN but found:" + p.current.Value)
		}
		p.nextToken()
		exp:=p.expr()
		p.nextToken()
		if p.current.Type != Lexer.RPAREN{
			return nil,errors.New("Expected RPAREN but found:" + p.current.Value)
		}
		var stamtments []shared.Node
		p.nextToken()
		if p.current.Type != Lexer.OpeningParen{
			return nil,errors.New("Expected OPENING PAREN but found:" + p.current.Value)
		}
		for{
			if p.current.Type == Lexer.ClosingParen {
				break
			}else {
				statment,err:=p.Statement()
				shared.Check(err)
				stamtments = append(stamtments,statment)
				p.nextToken()
			}
		}
		return IfNode{Expression: exp,statements: stamtments},nil
	default:
		return p.expr(),nil
	}

}

/*
	Factor
*/

func (p *Parser) factor() (shared.Node,error) {
	tok := p.current
	if tok.Type == Lexer.INT {
		val, _ := strconv.ParseFloat(tok.Value, 64)
		p.nextToken()
		return NumberNode{Value: val},nil
	}
	if tok.Type == Lexer.FLOAT {
		var valueStr = strings.ReplaceAll(tok.Value, ",", ".")
		val, _ := strconv.ParseFloat(valueStr, 64)
		p.nextToken()
		return NumberNode{Value: val},nil
	}
	if tok.Type == Lexer.IDENTIFIER {
		ident := tok.Value
		p.nextToken()
		return VariableAccessNode{Name: ident},nil
	}
	if tok.Type == Lexer.LPAREN {
		p.nextToken()
		exprNode := p.expr()
		if tok.Type == Lexer.RPAREN {
			return nil,errors.New("Expected RPAREN but found:" + p.current.Value)
		}
		p.nextToken()
		return exprNode,nil
	}
	if tok.Type == Lexer.BOOL {
		value := tok.Value
		p.nextToken()
		if value == "true" {
			return BooleanNode{Value: true},nil
		} else {
			return BooleanNode{Value: false},nil
		}
	}
	if tok.Type == Lexer.STRING {
		p.nextToken()
		return StringNode{Value: tok.Value},nil
	} else {
		return nil,errors.New("Syntax Error: Expected VALUE but found " + tok.Value)
	}
}

/*
* Term
 */
func (p *Parser) term() shared.Node {
	node,err := p.factor()
	shared.Check(err)
	for p.current.Type == Lexer.STAR || p.current.Type == Lexer.SLASH {
		op := p.current
		p.nextToken()
		fac,err:=p.factor()
		shared.Check(err)
		node = BinaryOperationNode{
			Left:     node,
			Operator: op.Value,
			Right:    fac,
		}
	}
	for p.current.Type == Lexer.GREATER || p.current.Type == Lexer.LESS {
		op := p.current
		p.nextToken()
		fac,err:=p.factor()
		shared.Check(err)
		node = ComparisonNode{
			Left:  node,
			Right: fac,
			Op:    op.Value,
		}
	}
	return node
}

/*
Expression
*/
func (p *Parser) expr() shared.Node {
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
	for p.current.Type == Lexer.AND {
		op := p.current
		p.nextToken()
		node = ComparisonNode{
			Left:  node,
			Right: p.term(),
			Op:    op.Value,
		}

	}
	return node
}
func (p *Parser) Parse() shared.Node {
	program := ProgramNode{}
	for {
		if p.current.Type == Lexer.EOF {
			break
		}
		stmt,err:=p.Statement()
		shared.Check(err)
		program.statements = append(program.statements, stmt)
	}
	return program
}
