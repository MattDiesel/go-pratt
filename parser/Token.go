package parser

import (
	"fmt"
)

type Token struct { // implements: IToken
	name string
	lbp Precedence
	parent IParser
}

	func (this Token) Name() string {
		return this.name
	}

	func (this Token) Lbp() Precedence {
		return this.lbp
	}

	func (this Token) Nud() (IValue, error) {
		return nil, NewParserError(fmt.Sprintf("Token %s cannot appear prefix.", this.Name()))
	}

	func (this Token) Led(left IValue) (IValue, error) {
		return nil, NewParserError(fmt.Sprintf("Token %s cannot appear infix.", this.Name()))
	}

	func (this Token) SetParser(p IParser) {
		this.parent = p
	}

	func (this Token) GetParser() IParser {
		return this.parent
	}

	func NewToken(n string, bp Precedence) Token {
		return Token{n, bp, nil}
	}