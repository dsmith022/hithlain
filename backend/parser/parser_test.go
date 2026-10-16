package parser

import "testing"

func TestNextToken(t *testing.T) {
	input := `You've got a friend in me {$friend + 1}[Woody][Buzz]`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TEXT, "You've got a friend in me "},
		{LVAL, "{"},
		{VALSIGN, "$"},
		{VALNAME, "friend"},
		{PLUS, "+"},
		{INT, "5"},
		{RVAL, "}"},
		{LLINK, "["},
		{LINKEDNODE, "Woody"},
		{RLINK, "]"},
		{LLINK, "["},
		{LINKEDNODE, "Buzz"},
		{RLINK, "]"},
	}

	l := NewParser(input)

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
