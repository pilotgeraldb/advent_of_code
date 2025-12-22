package main

import (
	"day_5_aoc_2025/state"
	"log"
	shared "shared_aoc_2025"
	"strings"
)

func main() {
	total := 0
	s := state.NewState()
	shared.StreamProcess("input.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
			line := si.GetLine()
			if strings.Contains(line, "-") {
				p := strings.Split(line, "-")
				r := state.NewRangeFromStrings(p[0], p[1], true)
				s.AddRange(r)
			}
		default:

		}

	}, shared.WithEOFCallback(func() {
		total = s.UniqueCount()
		log.Printf("total: %d", total)
	}))
}
