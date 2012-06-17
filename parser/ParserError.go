
package parser

import (
	"fmt"
)

type Position struct { // implements: IValue
	line int
	col int
}
	func (this *Position) ToString() string {
		return fmt.Sprintf("%i:%i", this.line, this.col)
	}


type ParserError struct { // implements: error
	Pos* Position
	Msg string
	Name string

	// NewParserError(string, string, Position) ParserError
	// Error() string
}

func NewParserError(msg string) *ParserError {
	return &ParserError{nil, msg, "ParserError"}
}

func (this *ParserError) Error() string {
	if this.Pos == nil {
		return fmt.Sprintf("%s: %s", this.Name, this.Msg)
	}

	return fmt.Sprintf(
		"%s: %s\n\t%s",
		this.Pos.ToString(), this.Name, this.Msg)
}
