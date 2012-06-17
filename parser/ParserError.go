package parser

import (
	"fmt"
)

type Position struct { // implements: IValue
	Line int
	Col  int
}

func (this *Position) ToString() string {
	return fmt.Sprintf("%v:%v", this.Line, this.Col)
}

type ParserError struct { // implements: error
	Pos *Position
	Msg string

	// NewParserError(string, string, Position) ParserError
	// Error() string
}

func NewParserError(msg string) *ParserError {
	return &ParserError{nil, msg}
}

func (this *ParserError) Error() string {
	if this.Pos == nil {
		return fmt.Sprintf("ParserError: %v", this.Msg)
	}

	return fmt.Sprintf(
		"%v: ParserError\n\t%v",
		this.Pos.ToString(), this.Msg)
}
