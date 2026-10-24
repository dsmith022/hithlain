package parser

import (
	"strings"
)

type Lexer struct {
	body         string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(body string) *Lexer {
	l := &Lexer{body: body}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.body) {
		l.ch = 0
	} else {
		l.ch = l.body[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer) NextToken() Token {
	var token Token

	switch l.ch {
	//Identifiers
	case '{':
		token = newToken(LVAL, l.ch)

	case '}':
		token = newToken(RVAL, l.ch)
	case '[':
		token = newToken(LLINK, l.ch)
	case ']':
		token = newToken(RLINK, l.ch)

	case '$':
		token = newToken(VALSIGN, l.ch)

		//Primitive types
	case 0:
		token.Literal = ""
		token.Type = EOF

		//Operators
	case '+':
		token = newToken(PLUS, l.ch)
	case '-':
		token = newToken(MINUS, l.ch)
	case '=':
		token = newToken(EQUALS, l.ch)

	default:
		if isLetter(l.ch) {
			token.Type = TEXT

			if l.position > 0 {
				switch l.body[l.position-1] {
				case '$':
					token.Type = VALNAME
				case '[':
					token.Type = LINKEDNODE
				}
			}

			token.Literal = l.readText()
			return token

		} else if isWhitespace(l.ch) {
			token.Literal = l.readWhitespace()
			token.Type = WHITESPACE
			return token
		} else if isDigit(l.ch) {
			token.Literal = l.readNumber()
			token.Type = INT
			return token
		} else {
			token = newToken(ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return token
}

func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}

func (l *Lexer) readText() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.body[position:l.position]
}

func isLetter(ch byte) bool {
	illegals := "${}[]+-="

	return !isDigit(ch) && !isWhitespace(ch) && ch != 0 && !strings.ContainsRune(illegals, rune(ch))
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.body[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) readWhitespace() string {
	position := l.position
	for isWhitespace(l.ch) {
		l.readChar()
	}
	return l.body[position:l.position]
}

/*
skipWhitespace() will need to be used only within {} or [] blocks appropriately.
will need to give more though to exactly how to handle these cases,
as this is not as simple as writing a normal interpreter because of our TEXT types.
we have to intelligently process whether spaces are to be preserved or not,
based on the TokenType.

func (p *Parser) skipWhitespace() {
	for l .ch == ' ' || l .ch == '\t' || l .ch == '\n' || l .ch == '\r' {
		l .readChar()
	}

}
*/
