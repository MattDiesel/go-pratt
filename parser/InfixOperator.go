package parser

type LedHandler func(a, b IValue) (IValue, error)

type InfixOperator struct { // implements: IToken
	Symbol   string
	Bp       Precedence
	Function LedHandler
	Parent   IParser
}

func (this *InfixOperator) Name() string {
	return this.Symbol
}

func (this *InfixOperator) Lbp() Precedence {
	return this.Bp
}

func (this *InfixOperator) Led(left IValue) (IValue, error) {
	right, err := this.Parent.expression(this.Bp)

	if err != nil {
		return nil, err
	}

	return this.Function(left, right)
}

func (this *InfixOperator) Nud() (IValue, error) {
	return nil, &ParserError{
		this.Parent.Lexer().Here(),
		"Infix operator cannot appear prefix"}
}

func (this *InfixOperator) SetParser(p IParser) {
	this.Parent = p
}

func (this *InfixOperator) GetParser() IParser {
	return this.Parent
}
