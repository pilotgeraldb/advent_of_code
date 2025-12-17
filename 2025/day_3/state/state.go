package state

import (
	"fmt"
	"math"
)

type State struct {
	arr   []int
	topK  []int
	count int
}

func NewState(count int) State {
	return State{
		topK:  make([]int, count),
		arr:   []int{},
		count: count,
	}
}

func (s State) Calculate() int {
	total := 0
	for i, v := range s.topK {
		t := int(float64(v) * math.Pow(10, float64(s.count-i-1)))
		total += t
	}
	fmt.Printf(" topK: %v | total: %d", s.topK, total)
	return total
}

func (s *State) Reset() {
	s.topK = make([]int, s.count)
	s.arr = []int{}
}

func (s *State) SetInputArray(arr []int) {
	s.arr = arr
}

func (s *State) TopK() {
	s.topK = make([]int, s.count)
	arr := s.arr

	for i := 0; i < s.count; i++ {
		largestIdx := -1
		largestVal := math.MinInt

		remaining := s.count - i - 1

		for j := 0; j < len(arr)-remaining; j++ {
			if arr[j] > largestVal {
				largestVal = arr[j]
				largestIdx = j
			}
		}

		if largestIdx == -1 {
			panic("not enough elements left to pick topK")
		}

		s.topK[i] = largestVal
		arr = arr[largestIdx+1:]
	}
}
