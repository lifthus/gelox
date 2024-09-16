package parser

import (
	"github.com/lifthus/gelox/ast"
	"github.com/lifthus/gelox/lexer"
	"github.com/lifthus/gelox/token"
)

// Parser parses the tokens from the given lexer and generates the abstract syntax tree for Gelox.
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Gelox parser with the given lexer.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
