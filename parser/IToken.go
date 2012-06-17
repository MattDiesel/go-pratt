package parser

type IToken interface {
	Name() string
	Lbp() Precedence
	Nud() (IValue, error)
	Led(IValue) (IValue, error)
	GetParser() IParser
	SetParser(p IParser)
}
