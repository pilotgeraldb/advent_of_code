package state

import "strconv"

type State struct {
	ranges []Range
}

func NewState() *State {
	return &State{
		ranges: []Range{},
	}
}

func (s *State) AddRange(r Range) {
	s.ranges = append(s.ranges, r)
}

func (s State) Ranges() *[]Range {
	return &s.ranges
}

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) Range {
	return Range{
		Start: start,
		End:   end,
	}
}

func NewRangeFromStrings(start, end string) Range {
	s, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}
	return Range{
		Start: s,
		End:   e,
	}
}
