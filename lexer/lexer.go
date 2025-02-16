package lexer

import "github.com/lifthus/gelox/token"

const (
	ASCII_NUL = 0
)

// Lexer tokenizes the input string.
//
// If Unicode should be supported, the reading method should be changed and the field ch should be changed to rune.
// Because multiple bytes can be assigned for a single character in Unicode.
type Lexer struct {
	// input is the primitive target string to be tokenized.
	input string
	// position is the cursor that points to the current byte.
	position int
	// ch is the current byte.
	ch byte
	// readposition is the cursor that points the byte to be read.
	readPosition int
}

// New creates a new Gelox lexer with the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next token and advances the cursors.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch   l.ch {
	case '=':
		if l.peekNextChar() == '=' {
			if l.peekNextNextChar() == '=' {
				ch := l.ch
				l.readChar()
				nch := l.ch
				l.readChar()
				literal := string(ch) + string(nch) + string(l.ch)
				tok = token.Token{Type: token.STR_EQ, Literal: literal}
			} else {
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.EQ, Literal: literal}
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekNextChar() == '=' {
			if l.peekNextNextChar() == '=' {
				ch := l.ch
				l.readChar()
				nch := l.ch
				l.readChar()
				literal := string(ch) + string(nch) + string(l.ch)
				tok = token.Token{Type: token.STR_NEQ, Literal: literal}
			} else {
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.NEQ, Literal: literal}
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ASCII_NUL:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.DOUBLE
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken creates a new token with the given type and character.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readChar reads the next character from the input string and advances the position pointers.
//
// It returns ASCII_NUL if there is no character to read.
func (l *Lexer) readChar() {
	if len(l.input) <= l.readPosition {
		l.ch = ASCII_NUL // Not read or EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// readIdentifier reads the identifier from the input string advancing the position pointers.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace skips the whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter checks if the given character can be recognized as a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if the given character can be recognized as a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readNumber reads the number from the input string advancing the position pointers.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// peekNextChar returns the readPosition th character without advancing the position pointers.
func (l *Lexer) peekNextChar() byte {
	return l.peekCharAt(0)
}

// peekNextNextChar returns readPositioin + 1 th character without advancing the position pointers.
func (l *Lexer) peekNextNextChar() byte {
	return l.peekCharAt(1)
}

// peekCharAt returns the readPosition + offset th character without advancing the position pointers.
func (l *Lexer) peekCharAt(offset int) byte {
	targetPosition := l.readPosition + offset
	if len(l.input) <= targetPosition {
		return ASCII_NUL
	}
	return l.input[targetPosition]
}
