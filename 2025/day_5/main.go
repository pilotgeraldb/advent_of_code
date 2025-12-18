package main

import (
	"day_5_aoc_2025/state"
	"log"
	shared "shared_aoc_2025"
	"strconv"
	"strings"
)

func main() {
	s := state.NewState()
	endOfRanges := false
	total := 0
	shared.StreamProcess("input.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
			line := si.GetLine()
			if !endOfRanges {
				if strings.Contains(line, "-") {
					p := strings.Split(line, "-")
					r := state.NewRangeFromStrings(p[0], p[1])
					s.AddRange(r)
				}
				if line == "" {
					endOfRanges = true
				}
			} else {
				val, err := strconv.Atoi(line)
				if err != nil {
					panic(err)
				}
				ranges := *s.Ranges()
				for _, r := range ranges {
					if val >= r.Start && val <= r.End {
						total++
						break
					}
				}
			}
		default:

		}

	}, shared.WithEOFCallback(func() {
		log.Printf("total: %d", total)
	}))
}
