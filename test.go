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

type MyLexer struct {
	Pos    parser.Position
	Ch     <-chan parser.IToken
	Parent parser.IParser
}

func (this MyLexer) Read() (parser.IToken, error) {
	t, ok := <-this.Ch

	this.Pos.Col += 1

	if !ok {
		return this.Parent.GetSymbol("(end)"), nil
	}

	t.SetParser(this.Parent)
	return t, nil
}

func (this MyLexer) Here() *parser.Position {
	return &this.Pos
}

func NewLexer(ch chan parser.IToken, p parser.IParser, input string) {
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
		} else if c == '*' {
			ch <- p.GetSymbol("*")
		} else if c == '/' {
			ch <- p.GetSymbol("/")
		} else {
			i, _ := strconv.Atoi(fmt.Sprint(c))
			ch <- parser.NewLiteral(NewInteger(i - 48))
		}
	}
	ch <- p.GetSymbol("(end)")

	close(ch)
}

type addOp struct { // implements IToken
	Parent parser.IParser
}

func (this *addOp) Name() string {
	return "+"
}

func (this *addOp) Lbp() parser.Precedence {
	return 10
}

func (this *addOp) Led(left parser.IValue) (parser.IValue, error) {
	right, err := this.Parent.Expression(10)

	if err != nil {
		return nil, err
	}

	ai, _ := left.(*Integer)
	bi, _ := right.(*Integer)

	return NewInteger(ai.value + bi.value), nil
}

func (this *addOp) Nud() (parser.IValue, error) {
	return this.Parent.Expression(30)
}

func (this *addOp) SetParser(p parser.IParser) {
	this.Parent = p
}

func (this *addOp) GetParser() parser.IParser {
	return this.Parent
}

type subOp struct { // implements IToken
	Parent parser.IParser
}

func (this *subOp) Name() string {
	return "-"
}

func (this *subOp) Lbp() parser.Precedence {
	return 10
}

func (this *subOp) Led(left parser.IValue) (parser.IValue, error) {
	right, err := this.Parent.Expression(10)

	if err != nil {
		return nil, err
	}

	ai, _ := left.(*Integer)
	bi, _ := right.(*Integer)

	return NewInteger(ai.value - bi.value), nil
}

func (this *subOp) Nud() (parser.IValue, error) {
	right, err := this.Parent.Expression(30)

	if err != nil {
		return nil, err
	}

	bi, _ := right.(*Integer)

	return NewInteger(-bi.value), nil
}

func (this *subOp) SetParser(p parser.IParser) {
	this.Parent = p
}

func (this *subOp) GetParser() parser.IParser {
	return this.Parent
}

func main() {
	p := parser.NewParser()

	mul := &parser.InfixOperator{"*", 20, func(a, b parser.IValue) (parser.IValue, error) {
		ai, _ := a.(*Integer)
		bi, _ := b.(*Integer)

		return NewInteger(ai.value * bi.value), nil
	}, nil}

	div := &parser.InfixOperator{"/", 20, func(a, b parser.IValue) (parser.IValue, error) {
		ai, _ := a.(*Integer)
		bi, _ := b.(*Integer)

		return NewInteger(ai.value / bi.value), nil
	}, nil}

	end := &parser.Token{"(end)", -1, nil}

	p.Add(&addOp{nil})
	p.Add(&subOp{nil})
	p.Add(mul)
	p.Add(div)
	p.Add(end)

	ch := make(chan parser.IToken)
	go NewLexer(ch, p, "5+4*-3")

	lexer := &MyLexer{parser.Position{1, 0}, ch, p}

	v, err := p.Parse(lexer)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v.ToString())
	}
}
