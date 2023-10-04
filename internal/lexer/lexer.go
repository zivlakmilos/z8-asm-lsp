package lexer

type Lexer struct {
	input        string
	position     int
	peekPosition int
	ch           byte
}

var tokenTypeMap = map[string]TokenType{
	"nop":  TokenTypeNop,
	"lda":  TokenTypeLdA,
	"ldb":  TokenTypeLdB,
	"sta":  TokenTypeStA,
	"stb":  TokenTypeStB,
	"seta": TokenTypeSetA,
	"setb": TokenTypeSetB,
	"not":  TokenTypeNot,
	"and":  TokenTypeAnd,
	"or":   TokenTypeOr,
	"xor":  TokenTypeXor,
	"nand": TokenTypeNand,
	"nor":  TokenTypeNor,
	"add":  TokenTypeAdd,
	"sub":  TokenTypeSub,
	"jmp":  TokenTypeJmp,
	"jmpz": TokenTypeJmpZ,
	"jmpc": TokenTypeJmpC,
	"hlt":  TokenTypeHlt,
}

func NewLexer(input string) *Lexer {
	lexer := Lexer{
		input:        input,
		position:     0,
		peekPosition: 0,
	}

	lexer.readChar()

	return &lexer
}

func (l *Lexer) Next() *Token {
	var tok *Token = nil

	l.skipWhitespaces()

	if l.ch == 0 {
		tok = NewToken(TokenTypeEof, "EOF")
	} else if l.isNumber(l.ch) && l.peek() == 'x' {
		literal := l.readHex()
		tok = NewToken(TokenTypeHex, literal)
	} else if l.ch == '0' && l.peek() == 'b' {
		literal := l.readBin()
		tok = NewToken(TokenTypeBin, literal)
	} else if l.isNumber(l.ch) {
		literal := l.readInt()
		tok = NewToken(TokenTypeInt, literal)
	} else if l.isLetter(l.ch) {
		literal := l.readLiteral()

		tokenType, ok := tokenTypeMap[literal]
		if !ok {
			tokenType = TokenTypeLiteral
		}

		tok = NewToken(tokenType, literal)
	}

	if tok == nil {
		tok = NewToken(TokenTypeIllegal, "")
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.peekPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.peekPosition]
	}

	l.position = l.peekPosition
	l.peekPosition++
}

func (l *Lexer) peek() byte {
	if l.peekPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.peekPosition]
	}
}

func (l *Lexer) readLiteral() string {
	startPos := l.position

	for l.isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPos:(l.position)]
}

func (l *Lexer) readInt() string {
	startPos := l.position

	for l.isNumber(l.ch) {
		l.readChar()
	}

	return l.input[startPos:(l.position)]
}

func (l *Lexer) readHex() string {
	startPos := l.position

	for l.isHex(l.ch) || (l.position-startPos == 1 && l.ch == 'x') {
		l.readChar()
	}

	return l.input[startPos:(l.position)]
}

func (l *Lexer) readBin() string {
	startPos := l.position

	for l.isBin(l.ch) || (l.position-startPos == 1 && l.ch == 'b') {
		l.readChar()
	}

	return l.input[startPos:(l.position)]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'A')
}

func (l *Lexer) isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) isBin(ch byte) bool {
	return ch == '0' || ch == '1'
}

func (l *Lexer) isHex(ch byte) bool {
	return (ch >= '0' && ch <= '9') ||
		(ch >= 'A' && ch <= 'F') ||
		(ch >= 'a' && ch <= 'f')
}
