package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LVAL    = "{"
	RVAL    = "}"
	LLINK   = "["
	RLINK   = "]"
	VALSIGN = "$"

	EOF     = ""
	TEXT    = "TEXT"
	NEWLINE = "\n"

	PLUS   = "+"
	MINUS  = "-"
	EQUALS = "="
)
