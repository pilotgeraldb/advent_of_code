package expression

import "fmt"

type OperandBuilder struct {
	runes []rune
}

func NewOperandBuilder() *OperandBuilder {
	return &OperandBuilder{
		runes: []rune{},
	}
}

func (ob *OperandBuilder) Append(r rune) {
	if r == ' ' || r == '\n' || r == '\r' {
		return
	}
	ob.runes = append(ob.runes, r)
}

func (ob OperandBuilder) String() string {
	return fmt.Sprintf("%s", string(ob.runes))
}

func (ob OperandBuilder) Build() *Operand {
	return NewOperandFromRunes(ob.runes)
}

func (ob *OperandBuilder) Reset() {
	ob.runes = []rune{}
}

func (ob OperandBuilder) IsEmpty() bool {
	hasRunes := len(ob.runes) == 0
	return hasRunes
}
