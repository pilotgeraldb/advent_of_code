package state

import (
	"day_2_aoc_2025/id"
	"strconv"
)

type State struct {
	Value       []rune
	Pairs       []id.Pair
	CurrentPair *id.Pair
	Result      int64
}

func NewState() State {
	return State{
		Value:       []rune{},
		Pairs:       []id.Pair{},
		CurrentPair: nil,
		Result:      0,
	}
}

func (s *State) Reset() {
	s.Value = []rune{}
	s.CurrentPair = nil
}

func (s State) GetValue() int64 {
	v, err := strconv.ParseInt(string(s.Value), 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func (s *State) Validate() {
	for i := range s.Pairs {
		s.Pairs[i].Validate()
	}
}

func (s *State) Sum() int64 {
	r := int64(0)
	for _, pair := range s.Pairs {
		// adding up all the invalid IDs
		r += pair.SumInvalid()
	}
	return r
}
