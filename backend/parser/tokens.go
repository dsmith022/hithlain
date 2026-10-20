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
	EOF        = ""
	WHITESPACE = "WHITESPACE"
	TEXT       = "TEXT"
	INT        = "INT"
	ILLEGAL    = "ILLEGAL"

	//Operators
	PLUS   = "+"
	MINUS  = "-"
	EQUALS = "="
)

/*Do we need these functions?
 Because we will not be parsing keywords (ie let, fn)
 in the way that Monkey does in the interpreter reference. TBD



var keywords = map[string]TokenType{}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return TEXT
}*/
