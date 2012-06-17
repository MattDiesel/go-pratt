package main

import (
	"./parser"
	"fmt"
	"strconv"
)

type Integer struct { // implements parser.IValue
	value int
}

func (this *Integer) ToString() string {
	return fmt.Sprint(this.value)
}

func NewInteger(i int) *Integer {
	return &Integer{i}
}

func NewLexer(p parser.IParser, ch chan<- parser.IToken, input string) {
	defer func() {
		if x := recover(); x != nil {
			close(ch)
		}
	}()

	for _, c := range input {
		if c == '+' {
			ch <- p.GetSymbol("+")
		} else if c == '-' {
			ch <- p.GetSymbol("-")
		} else {
			i, _ := strconv.Atoi(fmt.Sprint(c))
			ch <- parser.NewLiteral(NewInteger(i - 48))
		}
	}
	ch <- p.GetSymbol("(end)")

	close(ch)
}

func main() {
	p := parser.NewParser()

	add := parser.NewInfixOperator("+", 10, func(a, b parser.IValue) (parser.IValue, error) {
		ai, _ := a.(*Integer)
		bi, _ := b.(*Integer)

		return NewInteger(ai.value + bi.value), nil
	})

	sub := parser.NewInfixOperator("-", 10, func(a, b parser.IValue) (parser.IValue, error) {
		ai, _ := a.(*Integer)
		bi, _ := b.(*Integer)

		return NewInteger(ai.value - bi.value), nil
	})

	end := parser.NewToken("(end)", -1)

	p.Add(add)
	p.Add(sub)
	p.Add(end)

	ch := make(chan parser.IToken)
	go NewLexer(p, ch, "5+4-3")

	v, err := p.Parse(ch)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v.ToString())
	}
}
