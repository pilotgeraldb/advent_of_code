package main

import (
	"day_6_aoc_2025/expression"
	"day_6_aoc_2025/state"
	"log"
	shared "shared_aoc_2025"
)

func main() {
	s := state.New()
	shared.StreamProcess("input.txt", func(si shared.StreamInfo) {
		currentOperand := s.CurrentOperand()
		switch si.R {
		case '+', '-', '*', '/':
			exp := s.CurrentExpression()
			exp.SetOperation(expression.OperatorFromRune(si.R.Rune()))
			s.NextExpression()
		case ' ':
			AddToExpression(&s, si)
		default: // anything else
			currentOperand.Append(si.R.Rune())
		}
	}, shared.WithEOFCallback(func() {
		total := 0
		for _, exp := range s.Expressions() {
			total += exp.Evaluate()
		}
		log.Printf("total: %d", total)
	}), shared.WithNewlineCallback(func() {
		if !s.CurrentOperand().IsEmpty() {
			log.Println("not empty")
		}
		s.Reset()
	}), shared.WithCRCallback(func(si shared.StreamInfo) {
		AddToExpression(&s, si)
	}))
}

func AddToExpression(s *state.State, si shared.StreamInfo) {
	currentOperand := s.CurrentOperand()
	if currentOperand.IsEmpty() {
		return
	}
	if si.Line == 1 {
		exp := expression.NewExpression()
		op := currentOperand.Build()
		exp.AppendOperand(op)
		s.AddExpression(exp)
	} else {
		exp := s.CurrentExpression()
		op := currentOperand.Build()
		exp.AppendOperand(op)
	}
	s.NextExpression()
}

// 7098065460541 -- correct
