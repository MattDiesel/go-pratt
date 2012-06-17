package parser

type Literal struct { // implements: IToken
	value IValue
	parent IParser
}
	func (this *Literal) Name() string {
		return "(literal)"
	}

	func (this *Literal) Lbp() Precedence {
		return 0
	}

	func (this *Literal) Led(left IValue) (IValue, error) {
		return nil, NewParserError("Literal expression cannot appear infix.")
	}

	func (this *Literal) Nud() (IValue, error) {
		return this.value, nil
	}

	func (this *Literal) SetParser(p IParser) {
		this.parent = p
	}

	func (this *Literal) GetParser() IParser {
		return this.parent
	}

	func NewLiteral(val IValue) *Literal {
		return &Literal{val, nil}
	}