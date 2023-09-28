package lexer

type TokenType int

const (
	TokenTypeIllegal = iota
	TokenTypeEof
	TokenTypeInt
	TokenTypeHex
	TokenTypeBin
	TokenTypeLiteral
	TokenTypeNop
	TokenTypeLdA
	TokenTypeLdB
	TokenTypeStA
	TokenTypeStB
	TokenTypeSetA
	TokenTypeSetB
	TokenTypeNot
	TokenTypeAnd
	TokenTypeOr
	TokenTypeXor
	TokenTypeNand
	TokenTypeNor
	TokenTypeAdd
	TokenTypeSub
	TokenTypeJmp
	TokenTypeJmpZ
	TokenTypeJmpC
	TokenTypeHlt
)

type Token struct {
	tokenType TokenType
	literal   string
}

func NewToken(tokenType TokenType, literal string) *Token {
	token := Token{
		tokenType: tokenType,
		literal:   literal,
	}

	return &token
}
