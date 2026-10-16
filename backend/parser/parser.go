package parser

type Parser struct {
	body         string
	position     int
	readPosition int
	ch           byte
}

func NewParser(body string) *Parser {
	p := &Parser{body: body}
	p.readChar()
	return p
}

func (p *Parser) readChar() {
	if p.readPosition >= len(p.body) {
		p.ch = 0
	} else {
		p.ch = p.body[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition += 1

}

func (p *Parser) NextToken() Token {
	var token Token

	switch p.ch {
	//Identifiers
	case '{':
		token = newToken(LVAL, p.ch)

	case '}':
		token = newToken(RVAL, p.ch)
	case '[':
		token = newToken(LLINK, p.ch)
	case ']':
		token = newToken(RLINK, p.ch)

	case '$':
		token = newToken(VALSIGN, p.ch)

		//Primitive types
	case 0:
		token.Literal = ""
		token.Type = EOF

		//Operators
	case '+':
		token = newToken(PLUS, p.ch)
	case '-':
		token = newToken(MINUS, p.ch)
	case '=':
		token = newToken(EQUALS, p.ch)

	default:
		token.Literal = p.readIdentifier()
		token.Type = LookupIdent(token.Literal)

	}
	p.readChar()
	return token
}

func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}

func (p *Parser) readIdentifier() string {
	position := p.position
	for isText(p.ch) {
		p.readChar()
	}
	return p.body[position:p.position]
}

/*
Will be used to parse VAL or LINK names.
Nodes or Value Names must be either letters or underscores
*/
func isText(ch byte) bool {
	invalidChars := [...]byte{'[', ']', '{', '}', '$'}
	for _, v := range invalidChars {
		if ch == v {
			return false
		}

	}

	return true
}

/*
skipWhitespace() will need to be used only within {} or [] blocks appropriately.
will need to give more though to exactly how to handle these cases,
as this is not as simple as writing a normal interpreter because of our TEXT types.
we have to intelligently process whether spaces are to be preserved or not,
based on the TokenType.
*/
func (p *Parser) skipWhitespace() {
	for p.ch == ' ' || p.ch == '\t' || p.ch == '\n' || p.ch == '\r' {
		p.readChar()
	}

}
