package expression

import (
	"fmt"
	"log"
)

type Expression struct {
	left     *Operand
	right    *Operand
	operator Operator
	next     *Expression
}

func NewExpression() *Expression {
	return &Expression{
		left:     nil,
		right:    nil,
		operator: None,
		next:     nil,
	}
}

func (e *Expression) AppendOperand(o *Operand) {
	if e.left == nil {
		e.left = o
	} else if e.right == nil {
		e.right = o
	} else if e.next != nil {
		e.next.AppendOperand(o)
	} else {
		e.next = &Expression{
			left: o,
		}
	}
}

func (e *Expression) SetOperation(o Operator) {
	if e.operator == None {
		e.operator = o
	}
	if e.next != nil {
		e.next.SetOperation(o)
	}
}

func (e Expression) evaluateSingle() int {
	if e.left == nil {
		if e.operator == Add || e.operator == Subtract {
			e.left = &Operand{Value: 0}
		} else {
			e.left = &Operand{Value: 1}
		}
	}
	if e.right == nil {
		if e.operator == Add || e.operator == Subtract {
			e.right = &Operand{Value: 0}
		} else {
			e.right = &Operand{Value: 1}
		}
	}
	switch e.operator {
	case Add:
		return e.left.Value + e.right.Value
	case Subtract:
		return e.left.Value - e.right.Value
	case Multiply:
		return e.left.Value * e.right.Value
	case Divide:
		return e.left.Value / e.right.Value
	default:
		return 0
	}
}

func (e Expression) Evaluate() int {
	result := e.evaluateSingle()
	if e.next != nil {
		nextExp := *e.next
		nextExp.AppendOperand(NewOperandFromValue(result))
		if nextExp.operator == None {
			nextExp.operator = e.operator
		}
		log.Printf("evalute: `%s = %d`", e.String(), result)
		return nextExp.Evaluate()
	}
	log.Printf("evalute: `%s = %d`", e.String(), result)
	return result
}

func (e Expression) String() string {
	return fmt.Sprintf("%s %s %s", e.left, e.operator.String(), e.right)
}
