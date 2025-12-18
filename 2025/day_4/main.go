package main

import (
	"day_4_aoc_2025/state"
	shared "shared_aoc_2025"
)

const max_adjacent_rolls = 3
const row_length = 140
const char = '@'

func main() {
	s := state.NewState(row_length, max_adjacent_rolls)
	shared.StreamProcess("input.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
		default:
			s.Add(si.R == char)
		}
	}, shared.WithEOFCallback(func() {
		println(s.Evaluate())
	}))
}
