package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (

	//Syntax
	LVAL    = "{"
	RVAL    = "}"
	LLINK   = "["
	RLINK   = "]"
	VALSIGN = "$"

	//Identifiers
	SET  = "SET"
	IF   = "IF"
	ELSE = "ELSE"
	NEW  = "NEW"

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
	TIMES  = "*"
	DIV    = "/"
)

var keywords = map[string]TokenType{
	"set":  SET,
	"if":   IF,
	"else": ELSE,
	"new":  NEW,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return TEXT
}
