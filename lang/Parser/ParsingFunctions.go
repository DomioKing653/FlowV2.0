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
	var needArg = false
	p.nextToken()
	if p.current.Type != Lexer.IDENTIFIER {
		return nil, errors.New("expected identifier in function defiition")
	}
	var id = p.current.Value
	p.nextToken()
	if p.current.Type != Lexer.LPAREN {
		return nil, errors.New("expected left paren in function defiition")
	}
	p.nextToken()
	var arguments []string
	for {
		if p.current.Type == Lexer.EOF {
			return nil, errors.New("unexpected EOF in function definition")
		} else {
			if p.current.Type == Lexer.IDENTIFIER {
				arguments = append(arguments, p.current.Value)
				p.nextToken()
				if p.current.Type == Lexer.RPAREN {
					needArg = false
					break
				}
				if p.current.Type == Lexer.COMMA {
					needArg = true
					continue
				}
			} else {
				if !needArg {
					if p.current.Type == Lexer.RPAREN {
						break
					}
				} else {
					return nil, errors.New("error while parsing arguments")
				}
			}
		}
	}
	p.nextToken()
	var statments []shared.Node
	if p.current.Type != Lexer.OpeningParen {
		return nil, errors.New("expected openning paren")
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
	return FunctionNode{args: arguments, statments: statments, id: id}, nil
}
