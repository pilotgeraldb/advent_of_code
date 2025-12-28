package expression

import "fmt"

type Operator int

const (
	None Operator = iota
	Add
	Subtract
	Multiply
	Divide
)

func OperatorFromRune(r rune) Operator {
	switch r {
	case '+':
		return Add
	case '-':
		return Subtract
	case '*':
		return Multiply
	case '/':
		return Divide
	default:
		panic(fmt.Sprintf("operator not supported: %c", r))
	}
}

func (o Operator) String() string {
	switch o {
	case Add:
		return "+"
	case Subtract:
		return "-"
	case Multiply:
		return "x"
	case Divide:
		return "/"
	default:
		return "None"
	}
}
