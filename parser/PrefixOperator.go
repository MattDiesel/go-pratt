package parser

import (
	"fmt"
)

type PrefixOp func(a IValue) (IValue, error)

type PrefixOperator struct { // implements: IToken
	Symbol   string
	Bp       Precedence
	Function PrefixOp
	Parent   IParser
}

func (this *PrefixOperator) Name() string {
	return this.Symbol
}

func (this *PrefixOperator) Lbp() Precedence {
	return this.Bp
}

func (this *PrefixOperator) Led(left IValue) (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		fmt.Sprint("Prefix operator '%v' cannot appear infix", this.Name())}
}

func (this *PrefixOperator) Nud() (IValue, error) {
	right, err := this.Parent.Expression(this.Lbp())

	if err != nil {
		return nil, err
	}

	return this.Function(right)
}

func (this *PrefixOperator) SetParser(p IParser) {
	this.Parent = p
}

func (this *PrefixOperator) GetParser() IParser {
	return this.Parent
}
