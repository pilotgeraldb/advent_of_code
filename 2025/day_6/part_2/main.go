package main

import (
	//"day_6_part_2_aoc_2025/state"
	"log"
	shared "shared_aoc_2025"
)

func main() {
	//s := state.New()
	shared.StreamProcess("test.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '+', '-', '*', '/':
		case ' ':
		default:
		}
	}, shared.WithEOFCallback(func() {
		total := 0
		log.Printf("total: %d", total)
	}), shared.WithNewlineCallback(func(si shared.StreamInfo) {

	}), shared.WithCRCallback(func(si shared.StreamInfo) {

	}))
}

// 7098065460541 -- correct
