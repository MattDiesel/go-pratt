package parser

import (
	"fmt"
)

type Token struct { // implements: IToken
	Type   string
	Bp     Precedence
	Parent IParser
}

func (this *Token) Name() string {
	return this.Type
}

func (this *Token) Lbp() Precedence {
	return this.Bp
}

func (this *Token) Nud() (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		fmt.Sprintf("Token %s cannot appear prefix.", this.Name())}
}

func (this *Token) Led(left IValue) (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		fmt.Sprintf("Token %s cannot appear infix.", this.Name())}
}

func (this *Token) SetParser(p IParser) {
	this.Parent = p
}

func (this *Token) GetParser() IParser {
	return this.Parent
}
