package lexer

import (
	"testing"

	"yosaka/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if string(tok.Type) != string(tt.expectedType) {
			t.Fatalf("tests[%d] - type wrong, expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if token.TokenType(tok.Literal) != token.TokenType(tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong, expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
