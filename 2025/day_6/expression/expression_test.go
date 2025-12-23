package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateExpression(t *testing.T) {
	exp1 := NewExpression()
	exp1.AppendOperand(NewOperandFromValue(5))
	exp1.AppendOperand(NewOperandFromValue(5))
	exp1.SetOperation(Add)

	exp2 := NewExpression()
	exp2.AppendOperand(NewOperandFromValue(5))

	exp1.next = exp2

	actual := exp1.Evaluate()

	assert.Equal(t, 15, actual, "should be equal")
}
