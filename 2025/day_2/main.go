package main

import (
	"day_2_aoc_2025/id"
	"day_2_aoc_2025/state"
	"log"
	shared "shared_aoc_2025"
	"strconv"
)

func main() {
	s := state.State{}
	shared.StreamProcess("input.txt", func(r rune) {
		if s.CurrentPair == nil {
			s.CurrentPair = new(id.Pair)
		}
		switch r {
		case '\r':
			s.CurrentPair.Last, _ = strconv.ParseInt(string(s.Value), 10, 64)
		case '\n': // eof
			if s.CurrentPair != nil {
				s.Pairs = append(s.Pairs, *s.CurrentPair)
				s.Reset()
			}
			s.Validate()
			s.Result = s.Sum()
		case '-':
			s.CurrentPair.First, _ = strconv.ParseInt(string(s.Value), 10, 64)
			s.Value = []rune{}
		case ',':
			s.CurrentPair.Last, _ = strconv.ParseInt(string(s.Value), 10, 64)
			s.Pairs = append(s.Pairs, *s.CurrentPair)
			s.Reset()
		default:
			s.Value = append(s.Value, r)
		}
	})
	log.Printf("answer: %d", s.Result)
}
