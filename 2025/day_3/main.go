package main

import (
	"day_3_aoc_2025/state"
	"fmt"
	shared "shared_aoc_2025"
)

func main() {
	total := 0
	arr := []int{}
	s := state.NewState(12)

	shared.StreamProcess("input.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '\r':
		case '\n':
			s.SetInputArray(arr)
			s.TopK()
			total += s.Calculate()
			fmt.Printf(" | %d\n", total)
			s.Reset()
			arr = []int{}
		default:
			d := int(si.R - '0')
			fmt.Print(d)
			arr = append(arr, d)
		}
	})

	println(total)
}
