package lexer_test

import (
	"testing"

	"github.com/zivlakmilos/z8-asm-lsp/internal/lexer"
)

func TestLexer(t *testing.T) {
	input := `lda 0x01
ldb 0b00000010
add
nop
sta 3
hlt`

	expected := []lexer.Token{
		*lexer.NewToken(lexer.TokenTypeLdA, "lda"),
		*lexer.NewToken(lexer.TokenTypeHex, "0x01"),
		*lexer.NewToken(lexer.TokenTypeLdB, "ldb"),
		*lexer.NewToken(lexer.TokenTypeBin, "0b00000010"),
		*lexer.NewToken(lexer.TokenTypeAdd, "add"),
		*lexer.NewToken(lexer.TokenTypeNop, "nop"),
		*lexer.NewToken(lexer.TokenTypeStA, "sta"),
		*lexer.NewToken(lexer.TokenTypeInt, "3"),
		*lexer.NewToken(lexer.TokenTypeHlt, "hlt"),
	}

	lexer := lexer.NewLexer(input)

	for _, ex := range expected {
		token := lexer.Next()

		if token.TokenType != ex.TokenType {
			t.Fatalf("expected token type %v but got %v", ex.TokenType, token.TokenType)
		}

		if token.Literal != ex.Literal {
			t.Fatalf("expected literal %v but got %v", ex.Literal, token.Literal)
		}
	}
}
