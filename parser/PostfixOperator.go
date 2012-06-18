package parser

import (
	"fmt"
)

type PostfixOp func(a IValue) (IValue, error)

type PostfixOperator struct { // implements: IToken
	Symbol   string
	Bp       Precedence
	Function PostfixOp
	Parent   IParser
}

func (this *PostfixOperator) Name() string {
	return this.Symbol
}

func (this *PostfixOperator) Lbp() Precedence {
	return this.Bp
}

func (this *PostfixOperator) Led(left IValue) (IValue, error) {
	return this.Function(left)
}

func (this *PostfixOperator) Nud() (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		fmt.Sprint("Postfix operator '%v' cannot appear prefix.", this.Name())}
}

func (this *PostfixOperator) SetParser(p IParser) {
	this.Parent = p
}

func (this *PostfixOperator) GetParser() IParser {
	return this.Parent
}
