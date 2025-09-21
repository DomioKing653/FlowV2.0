package Parser

import (
	"errors"

	"Flow2.0/lang/Lexer"
	"Flow2.0/lang/shared"
)

/*
Println parsing function
*/
func (p *Parser) ParsePrintLn() (shared.Node, error) {
	p.nextToken()
	if p.current.Type != Lexer.LPAREN {
		return nil, errors.New("Expected LPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	exprNode := p.expr()
	if p.current.Type != Lexer.RPAREN {
		return nil, errors.New("Expected RPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	return PrintLnNode{exprNode}, nil
}

/*
If parsing function
*/
func (p *Parser) ParseIf() (shared.Node, error) {
	p.nextToken()
	if p.current.Type != Lexer.LPAREN {
		return nil, errors.New("Expected LPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	exprNode := p.expr()
	if p.current.Type != Lexer.RPAREN {
		return nil, errors.New("Expected RPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	var statments []shared.Node
	if p.current.Type != Lexer.OpeningParen {
		return nil, errors.New("Expected OPENING PAREN but found:" + p.current.Value)
	}
	p.nextToken()
	for {
		if p.current.Type == Lexer.ClosingParen {
			break
		} else {
			statment, err := p.Statement()
			shared.Check(err)
			statments = append(statments, statment)
		}
	}
	p.nextToken()
	return IfNode{Expression: exprNode, statements: statments}, nil
}

/*
While parsing function
*/
func (p *Parser) ParseWhile() (shared.Node, error) {
	p.nextToken()
	if p.current.Type != Lexer.LPAREN {
		return nil, errors.New("Expected LPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	exprNode := p.expr()
	if p.current.Type != Lexer.RPAREN {
		return nil, errors.New("Expected RPAREN but found:" + p.current.Value)
	}
	p.nextToken()
	var statments []shared.Node
	if p.current.Type != Lexer.OpeningParen {
		return nil, errors.New("Expected OPENING PAREN but found:" + p.current.Value)
	}
	p.nextToken()
	for {
		if p.current.Type == Lexer.ClosingParen {
			break
		}
		if p.current.Type == Lexer.EOF {
			return nil, errors.New("expected closing paren but found EOF")
		} else {
			statment, err := p.Statement()
			shared.Check(err)
			statments = append(statments, statment)
		}
	}
	p.nextToken()
	return WhileNode{Expression: exprNode, Statments: statments}, nil
}

/*
Loop parsing function
*/
func (p *Parser) ParseLoop() (shared.Node, error) {
	p.nextToken()
	var listNode []shared.Node
	if p.current.Type != Lexer.OpeningParen {
		return nil, errors.New("Expected OPPENING PAREN but found:" + p.current.Value)
	}
	p.nextToken()
	for {
		if p.current.Type == Lexer.ClosingParen {
			break
		} else {
			stmt, err := p.Statement()
			shared.Check(err)
			listNode = append(listNode, stmt)
		}
	}
	p.nextToken()
	return LoopNode{Nodes: listNode}, nil
}

/*
	Function Parsing
*/

func (p *Parser) ParseFunction() (shared.Node, error) {
	p.nextToken()
	if p.current.Type != Lexer.IDENTIFIER {
		return nil, errors.New("expected identifier in function definition")
	}
	id := p.current.Value
	p.nextToken()

	if p.current.Type != Lexer.LPAREN {
		return nil, errors.New("expected left paren in function definition")
	}
	p.nextToken()

	var arguments []string
	for p.current.Type != Lexer.RPAREN {
		if p.current.Type == Lexer.EOF {
			return nil, errors.New("unexpected EOF in function definition")
		}

		if p.current.Type != Lexer.IDENTIFIER {
			return nil, errors.New("expected identifier in argument list, got: " + p.current.Value)
		}
		arguments = append(arguments, p.current.Value)
		p.nextToken()

		if p.current.Type == Lexer.COMMA {
			p.nextToken()
		} else if p.current.Type != Lexer.RPAREN {
			return nil, errors.New("expected comma or closing paren after argument")
		}
	}

	p.nextToken()

	if p.current.Type != Lexer.OpeningParen {
		return nil, errors.New("expected opening paren for function body")
	}
	p.nextToken()

	var statements []shared.Node
	for p.current.Type != Lexer.ClosingParen {
		if p.current.Type == Lexer.EOF {
			return nil, errors.New("unexpected EOF while parsing function body")
		}

		stmt, err := p.Statement()
		shared.Check(err)
		statements = append(statements, stmt)
	}

	p.nextToken()

	return FunctionNode{args: arguments, statments: statements, id: id}, nil
}

/*
Parsing id
*/

func (p *Parser) ParseIdentifier() (shared.Node, error) {
	ident := p.current.Value
	p.nextToken()
	if p.current.Type == Lexer.EQUALS {
		p.nextToken()
		exprNode := p.expr()
		return VariableAssignNode{Name: ident, Value: exprNode}, nil
	}
	if p.current.Type == Lexer.LPAREN {
		p.nextToken()
	}
	return VariableAccessNode{Name: ident}, nil

}
