package parser

import ()

type IParser interface {
	Parse(ILexer) (IValue, error)
	expression(Precedence) (IValue, error)
	Add(IToken) error
	Step(IToken) error
	GetSymbol(string) IToken
	Lexer() ILexer
}
