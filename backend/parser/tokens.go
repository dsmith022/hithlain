package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (

	//Identifiers
	LVAL    = "{"
	RVAL    = "}"
	LLINK   = "["
	RLINK   = "]"
	VALSIGN = "$"

	//Value types
	VALNAME    = "VALNAME"
	LINKEDNODE = "LINKEDNODE"

	//Primitive types
	EOF     = ""
	TEXT    = "TEXT"
	INT     = "INT"
	ILLEGAL = "ILLEGAL"
	IDENT   = "IDENT"

	//Operators
	PLUS   = "+"
	MINUS  = "-"
	EQUALS = "="
)

// Placeholder for later
var keywords = map[string]TokenType{}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return TEXT
}
