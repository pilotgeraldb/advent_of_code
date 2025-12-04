package main

import (
	shared "shared_aoc_2025"
)

func main() {
	shared.StreamProcess("test.txt", func(r rune) {
		switch r {
		case '\r':
		case '\n':
		default:
		}
	})
}
