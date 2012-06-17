package parser

import ()

type IParser interface {
	Parse(<-chan IToken) (IValue, error)
	expression(Precedence) (IValue, error)
	Add(IToken) error
	Step(IToken) error
	GetSymbol(string) IToken
}
