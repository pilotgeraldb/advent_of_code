package state

import "log"

type State struct {
	arr         []bool
	rowLength   int
	maxAdjacent int
	removed     []bool
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
	arr := *s.Value()
	s.removed = make([]bool, len(arr))
	max := 5_000_000
	iteration := 0
	total := 0
	for {
		hasRemoval := false
		for i, v := range arr {
			if s.removed[i] {
				continue
			}
			if v {
				a := s.Adjacent(i)
				c := a.Count()
				if c <= s.maxAdjacent {
					//add to removal index
					s.removed[i] = true
					hasRemoval = true
					total++
				}
			}
		}

		if !hasRemoval {
			return total
		}

		iteration++
		if iteration > max {
			log.Fatalf("exceeded max iterations of %d", max)
		}
	}
}

func (s State) Adjacent(idx int) AdjacentResult {
	result := AdjacentResult{}

	//north
	if idx-s.rowLength >= 0 {
		if !s.removed[idx-s.rowLength] {
			result.North = s.arr[idx-s.rowLength]
		}
	}
	//south
	if idx+s.rowLength < len(s.arr) && idx+s.rowLength < len(s.removed) {
		if !s.removed[idx+s.rowLength] {
			result.South = s.arr[idx+s.rowLength]
		}
	}
	//west
	if idx%s.rowLength != 0 {
		if !s.removed[idx-1] {
			result.West = s.arr[idx-1]
		}
	}
	//east
	if (idx+1)%s.rowLength != 0 {
		if !s.removed[idx+1] {
			result.East = s.arr[idx+1]
		}
	}
	//northwest
	if idx-s.rowLength >= 0 && idx%s.rowLength != 0 {
		if !s.removed[idx-s.rowLength-1] {
			result.NorthWest = s.arr[idx-s.rowLength-1]
		}
	}
	//northeast
	if idx-s.rowLength >= 0 && (idx+1)%s.rowLength != 0 {
		if !s.removed[idx-s.rowLength+1] {
			result.NorthEast = s.arr[idx-s.rowLength+1]
		}
	}
	//southwest
	if idx+s.rowLength < len(s.arr) && idx%s.rowLength != 0 {
		if !s.removed[idx+s.rowLength-1] {
			result.SouthWest = s.arr[idx+s.rowLength-1]
		}
	}
	//southeast
	if idx+s.rowLength < len(s.arr) && (idx+1)%s.rowLength != 0 {
		if !s.removed[idx+s.rowLength+1] {
			result.SouthEast = s.arr[idx+s.rowLength+1]
		}
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
