package main

import (
	shared "shared_aoc_2025"
)

func main() {
	//s := state.NewState()
	shared.StreamProcess("test.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
		default:
		}
	})
}
