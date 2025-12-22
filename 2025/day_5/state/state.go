package state

import "sort"

type State struct {
	ranges []Range
}

func NewState() *State {
	return &State{
		ranges: []Range{},
	}
}

func (s *State) AddRange(r Range) {
	s.ranges = append(s.ranges, r)
}

func (s State) MergeAll() []Range {
	if len(s.ranges) == 0 {
		return nil
	}

	//avoid changing the original slice on the struct
	ranges := make([]Range, len(s.ranges))
	copy(ranges, s.ranges)

	//sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	// preallocate the merged slice to the same length as the original to avoid the use of 'append'
	// and extra allocations then use count to track the number of merged ranges "actually" filled
	merged := make([]Range, len(ranges))
	merged[0] = ranges[0]
	count := 1

	for _, r := range ranges[1:] {
		// get a pointer to the last merged range
		// we use this to compare current to this
		// one and see if they overlap/adjacent
		// last represents the most recently added merged range
		last := &merged[count-1]

		// determine if the end/start should be treated as "touching"?
		// if either range is inclusive then the endpoint counts as part of the range
		// can merge even if "r.Start == last.End + 1"
		// if both ranges are exclusive, adjacency = 0 because endpoints should not be counted
		adjacency := 0
		if last.inclusive || r.inclusive {
			adjacency = 1
		}

		if r.Start <= last.End+adjacency { // check the current range overlaps or is close enough to merge with the "last" merged range
			if r.End > last.End {
				// extend the end of the last merged range if the current range goes further.
				last.End = r.End
			}
			// set merged range as inclusive if either was inclusive
			last.inclusive = last.inclusive || r.inclusive
		} else {
			// if the current range does not overlap or touch (adjacent) the "last" merged range then
			// we start a new merged range and increment the filled count.
			merged[count] = r
			count++
		}
	}

	return merged[:count]
}

func (s State) UniqueCount() int {
	if len(s.ranges) == 0 {
		return 0
	}

	merged := s.MergeAll()

	total := 0
	for _, r := range merged {
		if r.inclusive {
			total += r.End - r.Start + 1
		} else {
			total += r.End - r.Start
		}
	}

	return total
}

//345755049374932
