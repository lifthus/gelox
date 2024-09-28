# token

token package defines the token types for the lexer and parser of gelox.

it provides LookupIdent func that checks the given ident and returns the reserved keyword token or IDENT token based on the ident.

```go
ident := token.LookupIdent("lifthus")
reserved := token.LookupIdent("if")
```
