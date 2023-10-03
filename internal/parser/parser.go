package parser

import (
	"fmt"

	"github.com/zivlakmilos/z8-asm-lsp/internal/ast"
	"github.com/zivlakmilos/z8-asm-lsp/internal/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken *lexer.Token
	nextToken    *lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := Parser{
		lexer: lexer,
	}

	parser.parseNextToken()
	parser.parseNextToken()

	return &parser
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := ast.NewProgram()

	for p.currentToken != nil {
		if p.currentToken.TokenType == lexer.TokenTypeEof {
			break
		}

		if !p.isInstruction(p.currentToken) {
			return nil, fmt.Errorf("parser error: expected instruction but got %v", p.currentToken.Literal)
		}

		if p.hasArgument(p.currentToken) {
			if p.nextToken != nil && p.isArgument(p.nextToken) {
				program.AddInstruction(ast.NewInstruction(p.currentToken, p.nextToken))
				p.parseNextToken()
			} else {
				return nil, fmt.Errorf("parser error: expected argument for instruction %s", p.currentToken.Literal)
			}
		} else {
			program.AddInstruction(ast.NewInstruction(p.currentToken, nil))
		}

		p.parseNextToken()
	}

	return program, nil
}

func (p *Parser) parseNextToken() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.Next()
}

func (p *Parser) hasArgument(token *lexer.Token) bool {
	switch token.TokenType {
	case lexer.TokenTypeLdA,
		lexer.TokenTypeLdB,
		lexer.TokenTypeStA,
		lexer.TokenTypeStB,
		lexer.TokenTypeSetA,
		lexer.TokenTypeSetB,
		lexer.TokenTypeJmp,
		lexer.TokenTypeJmpZ,
		lexer.TokenTypeJmpC:
		return true
	}

	return false
}

func (p *Parser) isArgument(token *lexer.Token) bool {
	switch token.TokenType {
	case lexer.TokenTypeInt,
		lexer.TokenTypeHex,
		lexer.TokenTypeBin:
		return true
	}

	return false
}

func (p *Parser) isInstruction(token *lexer.Token) bool {
	switch token.TokenType {
	case lexer.TokenTypeNop,
		lexer.TokenTypeLdA,
		lexer.TokenTypeLdB,
		lexer.TokenTypeStA,
		lexer.TokenTypeStB,
		lexer.TokenTypeSetA,
		lexer.TokenTypeSetB,
		lexer.TokenTypeNot,
		lexer.TokenTypeAnd,
		lexer.TokenTypeOr,
		lexer.TokenTypeXor,
		lexer.TokenTypeNand,
		lexer.TokenTypeNor,
		lexer.TokenTypeAdd,
		lexer.TokenTypeSub,
		lexer.TokenTypeJmp,
		lexer.TokenTypeJmpZ,
		lexer.TokenTypeJmpC,
		lexer.TokenTypeHlt:
		return true
	}

	return false
}
