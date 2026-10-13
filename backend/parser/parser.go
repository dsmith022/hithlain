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
	case '{':
		token = newToken(LVAL, p.ch)

	case '}':
		token = newToken(RVAL, p.ch)
	case '[':
		token = newToken(LLINK, p.ch)
	case ']':
		token = newToken(RLINK, p.ch)

	case 0:
		token.Literal = ""
		token.Type = EOF

	case '+':
		token = newToken(PLUS, p.ch)
	case '-':
		token = newToken(MINUS, p.ch)
	case '=':
		token = newToken(TEXT, p.ch)

	default:
		newToken(TEXT, p.ch)

	}
	p.readChar()
	return token
}

func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}
