package main

import (
	"log"
	shared "shared_aoc_2025"
)

func main() {
	shared.StreamProcess("test.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
		default:

		}

	}, shared.WithEOFCallback(func() {
		log.Print("end of file")
	}))
}
