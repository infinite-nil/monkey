package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers / Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSING = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREM    = "("
	RPAREM    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Kaywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)