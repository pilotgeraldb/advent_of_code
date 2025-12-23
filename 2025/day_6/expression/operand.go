package expression

import "fmt"

type Operand struct {
	Value int
}

func NewOperand() Operand {
	return Operand{
		Value: 0,
	}
}

func NewOperandFromRunes(runes []rune) *Operand {
	num := 0
	for _, r := range runes {
		num = num*10 + int(r-'0')
	}
	return &Operand{
		Value: num,
	}
}

func NewOperandFromValue(val int) *Operand {
	return &Operand{
		Value: val,
	}
}

func (o Operand) String() string {
	return fmt.Sprintf("%d", o.Value)
}
