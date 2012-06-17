package parser

type ILexer interface {
	Read() (IToken, error)
	Here() *Position
}
