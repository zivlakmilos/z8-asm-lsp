package parser_test

import (
	"testing"

	"github.com/zivlakmilos/z8-asm-lsp/internal/lexer"
	"github.com/zivlakmilos/z8-asm-lsp/internal/parser"
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
	parser := parser.NewParser(lexer)

	ast, err := parser.ParseProgram()
	if err != nil {
		t.Fatalf("program parsing failed with %v", err)
		return
	}

	i := 0
	for _, inst := range ast.Instructions {
		ex := expected[i]

		if ex.Literal != inst.TokenLiteral() {
			t.Fatalf("expected token literal %v but got %v", ex.Literal, inst.TokenLiteral())
		}

		i++
		if inst.Arg() != nil {
			arg := inst.Arg()
			ex := expected[i]
			if ex.Literal != arg.TokenLiteral() {
				t.Fatalf("expected token literal %v but got %v", ex.Literal, arg.TokenLiteral())
			}

			i++
		}
	}
}
