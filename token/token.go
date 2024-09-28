package token

type Token struct {
	Type    TokenType
	Literal string
}

// TokenType indicates the type of token.
// Allowed token types are predefined below.
//
// To improve the performance, the token type could be defined as int or byte.
type TokenType string

// Tokens for the parsing process.
const (

	// ILLEGAL indicates the tokens that are not recognized.
	ILLEGAL = "ILLEGAL"
	// EOF lets the parser know when to stop.
	// In other words, it indicates the end of the input.
	EOF = "EOF"
)

// Tokens for the literal things.
const (
	// IDENT indicates identifiers.
	//
	// e.g. abc, lifthus, gelox1, a_b_c, ...
	IDENT = "IDENT"
	// DOUBLE indicates the double-precision floating-point number literals.
	// In default, the Number type is double.
	//
	// e.g. 1.0, 3.14, 42, ...
	DOUBLE = "DOUBLE"
	// BIGINT indicates the 64-bit integer literals.
	//
	// e.g. 1n, 42n, ...
	BIGINT = "BIGINT"
)

// Operators
const (
	// ASSIGN assigns right-hand side value to left-hand side identifier.
	//
	// e.g. let x = 42;
	ASSIGN = "="
	// PLUS adds left-hand side value and right-hand side value.
	//
	// e.g. 4 + 2
	PLUS = "+"
	// MINUS can be used as infix operator and prefix operator.
	// - infix: subtracts the right-hand side value from the left-hand side value.
	// - prefix: negates the right-hand side value.
	//
	// e.g. 4 - 2
	MINUS = "-"
	// BANG negates the right-hand side value generating a boolean value.
	//
	// e.g. !true
	BANG = "!"
	// ASTERISK multiplies left-hand side value and right-hand side value.
	//
	// e.g. 4 * 2
	ASTERISK = "*"
	// SLASH divides left-hand side value by right-hand side value.
	//
	// e.g. 4 / 2
	SLASH = "/"
	// REMAINDER calculates the remainder of the division of left-hand side value by right-hand side value.
	//
	// e.g. 4 % 2
	REMAINDER = "%"

	// LT generates a boolean value indicating whether the left-hand side value is less than the right-hand side value.
	//
	// e.g. 4 < 2
	LT = "<"
	// GT generates a boolean value indicating whether the left-hand side value is greater than the right-hand side value.
	//
	// e.g. 4 > 2
	GT = ">"

	// EQ generates a boolean value indicating whether the left-hand side value is equal to the right-hand side value.
	// It weakly compares the values.
	//
	// e.g. 4 == 2
	EQ = "=="
	// NEQ generates a boolean value indicating whether the left-hand side value is not equal to the right-hand side value.
	// It weakly compares the values.
	//
	// e.g. 4 != 2
	NEQ = "!="
	// STR_EQ generates a boolean value indicating whether the left-hand side value is equal to the right-hand side value.
	// It strictly compares the values.
	//
	// e.g. 4 === 2
	STR_EQ = "==="
	// STR_NEQ generates a boolean value indicating whether the left-hand side value is not equal to the right-hand side value.
	// It strictly compares the values.
	//
	// e.g. 4 !== 2
	STR_NEQ = "!=="
)

// Delimiters
const (
	// COMMA separates specific tokens.
	//
	// e.g. let x = 1, y = 2; [1, 2, 3]
	COMMA = ","
	// SEMICOLON indicates the end of a statement.
	//
	// e.g. let x = 1; let y = 2; for (let i = 0; i < 10; i++) { ... }
	SEMICOLON = ";"
)

// Grouping
const (
	// LPAREN literally indicates the left parenthesis.
	LPAREN = "("
	// RPAREN literally indicates the right parenthesis.
	RPAREN = ")"
	// LBRACE literally indicates the left brace.
	LBRACE = "{"
	// RBRACE literally indicates the right brace.
	RBRACE = "}"
)

// Keywords
// must be registered in the keywords map below.
const (
	// FUNCTION indicates the reserved function keyword.
	//
	// e.g. function add(x, y) { return x + y; }
	FUNCTION = "FUNCTION"
	// LET indicates the reserved let keyword.
	//
	// e.g. let x = 42;
	LET = "LET"
	// TRUE indicates the boolean true literal.
	//
	// e.g. if (true) { ... }
	TRUE = "TRUE"
	// FALSE indicates the boolean false literal.
	//
	// e.g. if (false) { ... }
	FALSE = "FALSE"
	// IF indicates the reserved if keyword.
	//
	// e.g. if (x < 10) { ... }
	IF = "IF"
	// ELSE indicates the reserved else keyword.
	//
	// e.g. if (x < 10) { ... } else { ... }
	ELSE = "ELSE"
	// RETURN indicates the reserved return keyword.
	//
	// e.g. return 42;
	RETURN = "RETURN"
)

// keywords maps the keyword literals to the corresponding token types.
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

// LookupIdent checks whether the given ident is a reserved keyword.
// If it is, it returns the corresponding keyword token type,
// otherwise, it returns the IDENT token type, which indicates the developer-defined identifier.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
