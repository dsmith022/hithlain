package parser

import "testing"

func TestNextToken(t *testing.T) {
	input := `You've got a friend in me.{$friend+1}[Woody][Buzz]`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TEXT, "You've"},
		{WHITESPACE, " "},
		{TEXT, "got"},
		{WHITESPACE, " "},
		{TEXT, "a"},
		{WHITESPACE, " "},
		{TEXT, "friend"},
		{WHITESPACE, " "},
		{TEXT, "in"},
		{WHITESPACE, " "},
		{TEXT, "me."},
		{LVAL, "{"},
		{VALSIGN, "$"},
		{VALNAME, "friend"},
		{PLUS, "+"},
		{INT, "1"},
		{RVAL, "}"},
		{LLINK, "["},
		{LINKEDNODE, "Woody"},
		{RLINK, "]"},
		{LLINK, "["},
		{LINKEDNODE, "Buzz"},
		{RLINK, "]"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

	}

}
