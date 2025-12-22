package state

import (
	"strconv"
)

type Range struct {
	Start     int
	End       int
	inclusive bool
}

func NewRange(start, end int, inclusive bool) Range {
	r := Range{
		Start:     start,
		End:       end,
		inclusive: inclusive,
	}

	return r
}

func NewRangeFromStrings(start, end string, inclusive bool) Range {
	s, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}
	return NewRange(s, e, inclusive)
}
