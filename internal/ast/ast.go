package ast

import (
	"fmt"

	"github.com/zivlakmilos/z8-asm-lsp/internal/lexer"
)

type Node interface {
	String() string
	TokenLiteral() string
}

type Number struct {
	token *lexer.Token
	value int
}

type Instruction struct {
	token *lexer.Token
	arg   *Number
}

type Program struct {
	Instructions []Node
}

func (p *Program) String() string {
	result := "["

	for _, inst := range p.Instructions {
		result += fmt.Sprintf(" %v, ", inst)
	}
	result += "]"

	return result
}

func NewNumber(token *lexer.Token) *Number {
	return &Number{
		token: token,
		value: 0, // TODO: Change to converted value from token literal
	}
}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.TokenLiteral())
}

func (n *Number) TokenLiteral() string {
	return n.token.Literal
}

func NewInstruction(token *lexer.Token, arg *lexer.Token) *Instruction {
	inst := Instruction{
		token: token,
		arg:   nil,
	}

	if arg != nil {
		inst.arg = NewNumber(token)
	}

	return &inst
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%v", i.TokenLiteral())
}

func (i *Instruction) TokenLiteral() string {
	return i.token.Literal
}
