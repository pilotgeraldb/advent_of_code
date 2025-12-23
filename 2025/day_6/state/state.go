package state

import "day_6_aoc_2025/expression"

type State struct {
	expressions    []*expression.Expression
	currentOperand *expression.OperandBuilder
	opIdx          int
}

func New() State {
	return State{
		expressions:    []*expression.Expression{},
		currentOperand: expression.NewOperandBuilder(),
		opIdx:          0,
	}
}

func (s *State) AddExpression(e *expression.Expression) {
	s.expressions = append(s.expressions, e)
}

func (s *State) Expressions() []*expression.Expression {
	return s.expressions
}

func (s *State) CurrentExpression() *expression.Expression {
	return s.expressions[s.opIdx]
}

func (s State) CurrentOperand() *expression.OperandBuilder {
	return s.currentOperand
}

func (s *State) ResetAll() {
	s.currentOperand.Reset()
	s.opIdx = 0
	s.expressions = []*expression.Expression{}
}

func (s *State) Reset() {
	s.currentOperand.Reset()
	s.opIdx = 0
}

func (s *State) NextExpression() {
	s.currentOperand.Reset()
	s.opIdx++
	if s.opIdx >= len(s.expressions) {
		s.opIdx = 0
	}
}
