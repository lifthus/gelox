package ast

import "github.com/lifthus/gelox/token"

// Node is a component of the abstract syntax tree of Gelox.
type Node interface {
	TokenLiteral() string
}

// Statement is a node that doesn't generate any value.
type Statement interface {
	Node
	statementNode()
}

// Expression is a node that generates a value.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the abstract syntax tree in Gelox.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement does assign a expression to a identifier.
//
// every valid program in Gelox is composed of a series of statements.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
