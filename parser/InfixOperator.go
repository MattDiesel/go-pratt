package parser

import (
	"fmt"
)

type InfixOp func(a, b IValue) (IValue, error)

type InfixOperator struct { // implements: IToken
	Symbol   string
	Bp       Precedence
	Function InfixOp
	Parent   IParser
}

func (this *InfixOperator) Name() string {
	return this.Symbol
}

func (this *InfixOperator) Lbp() Precedence {
	return this.Bp
}

func (this *InfixOperator) Led(left IValue) (IValue, error) {
	right, err := this.Parent.Expression(this.Bp)

	if err != nil {
		return nil, err
	}

	return this.Function(left, right)
}

func (this *InfixOperator) Nud() (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		fmt.Sprint("Infix operator '%v' cannot appear prefix.", this.Name())}
}

func (this *InfixOperator) SetParser(p IParser) {
	this.Parent = p
}

func (this *InfixOperator) GetParser() IParser {
	return this.Parent
}
