package parser

type LedHandler func(a, b IValue) (IValue, error)

type InfixOperator struct { // implements: IToken
	name string
	lbp Precedence
	function LedHandler
	parent IParser
}
	func (this InfixOperator) Name() string {
		return this.name
	}

	func (this InfixOperator) Lbp() Precedence {
		return this.lbp
	}

	func (this InfixOperator) Led(left IValue) (IValue, error) {
		right, err := this.parent.expression( this.lbp )

		if err != nil {
			return nil, err
		}

		return this.function(left, right)
	}

	func (this InfixOperator) Nud() (IValue, error) {
		return nil, NewParserError("Infix operator cannot appear prefix")
	}

	func (this InfixOperator) SetParser(p IParser) {
		this.parent = p
	}

	func (this InfixOperator) GetParser() IParser {
		return this.parent
	}

	func NewInfixOperator(n string, bp Precedence, f LedHandler) InfixOperator {
		return InfixOperator{n, bp, f, nil}
	}