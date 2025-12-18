package state

import "log"

type State struct {
	arr         []bool
	rowLength   int
	maxAdjacent int
}

func NewState(rowLength int, maxAdjacent int) *State {
	return &State{
		arr:         make([]bool, 0),
		rowLength:   rowLength,
		maxAdjacent: maxAdjacent,
	}
}

func (s *State) Add(v bool) {
	s.arr = append(s.arr, v)
}

func (s *State) Value() *[]bool {
	return &s.arr
}

func (s State) Evaluate() int {
	total := 0
	arr := *s.Value()
	for i, v := range arr {
		if v {
			a := s.Adjacent(i)
			c := a.Count()
			if c <= s.maxAdjacent {
				log.Printf("counting %d because c = %d", i+1, c)
				total++
			} else {
				log.Printf("not counting %d because c = %d", i+1, c)
			}
		} else {
			log.Printf("not counting %d because value is false", i+1)
		}
	}
	return total
}

func (s State) Adjacent(idx int) AdjacentResult {
	result := AdjacentResult{}

	//north
	if idx-s.rowLength >= 0 {
		result.North = s.arr[idx-s.rowLength]
	}
	//south
	if idx+s.rowLength < len(s.arr) {
		result.South = s.arr[idx+s.rowLength]
	}
	//west
	if idx%s.rowLength != 0 {
		result.West = s.arr[idx-1]
	}
	//east
	if (idx+1)%s.rowLength != 0 {
		result.East = s.arr[idx+1]
	}
	//northwest
	if idx-s.rowLength >= 0 && idx%s.rowLength != 0 {
		result.NorthWest = s.arr[idx-s.rowLength-1]
	}
	//northeast
	if idx-s.rowLength >= 0 && (idx+1)%s.rowLength != 0 {
		result.NorthEast = s.arr[idx-s.rowLength+1]
	}
	//southwest
	if idx+s.rowLength < len(s.arr) && idx%s.rowLength != 0 {
		result.SouthWest = s.arr[idx+s.rowLength-1]
	}
	//southeast
	if idx+s.rowLength < len(s.arr) && (idx+1)%s.rowLength != 0 {
		result.SouthEast = s.arr[idx+s.rowLength+1]
	}

	return result
}

type AdjacentResult struct {
	North     bool
	NorthEast bool
	NorthWest bool

	South     bool
	SouthEast bool
	SouthWest bool

	West bool
	East bool
}

func (ar AdjacentResult) Count() int {
	count := 0
	if ar.North {
		count++
	}
	if ar.NorthEast {
		count++
	}
	if ar.NorthWest {
		count++
	}
	if ar.South {
		count++
	}
	if ar.SouthEast {
		count++
	}
	if ar.SouthWest {
		count++
	}
	if ar.West {
		count++
	}
	if ar.East {
		count++
	}
	return count
}
