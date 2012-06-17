package parser

import (
	"fmt"
)

type Parser struct { // implements: IParser
	token IToken
	Lexer <-chan IToken
	Symbols map[string]IToken
}

func NewParser() *Parser {
	ret := new(Parser)
	ret.token = NewToken("(error)", 0)
	ret.Symbols = make(map[string]IToken)
	return ret
}

func (this *Parser) Parse(lex <-chan IToken) (IValue, error) {
	this.Lexer = lex

	if err := this.Step(nil); err != nil {
		return nil, err
	}

	return this.expression(0)
}

func (this *Parser) expression(rbp Precedence) (IValue, error) {
	t := this.token

	if err := this.Step( nil ); err != nil {
		return nil, err
	}

	var left IValue
	var err error

	// Handle first token of expression.
	if left,err = t.Nud(); err != nil {
		return nil, err
	}

	// Parser loop. Handles precedence.
	for rbp <= this.token.Lbp() {
		t = this.token

		if err = this.Step( nil ); err != nil {
			return nil, err
		}

		if left,err = t.Led(left); err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (this *Parser) Add(t IToken) error {
	// TODO: Check if name already exists.

	t.SetParser(this)
	this.Symbols[t.Name()] = t

	return nil
}

func (this *Parser) Step(t IToken) error {
	if t != nil {
		if this.token.Name() != t.Name() {
			return NewParserError(
				fmt.Sprintf("Expected token type %s, got %s.", t.Name(), this.token.Name()))
		}
	}

	t, ok := <-this.Lexer

	if !ok {
		this.token = this.Symbols["(end)"]
	} else {
		this.token = t
	}

	return nil
}

func (this *Parser) GetSymbol(n string) IToken {
	return this.Symbols[n]
}
