package id

import "strconv"

type Pair struct {
	First   int64
	Last    int64
	invalid []int64
}

func (id *Pair) Validate() {
	r := id.rangeInts()
	for _, n := range r {
		if !validInt64(n) {
			id.invalid = append(id.invalid, n)
		}
	}
}

func (id Pair) rangeInts() []int64 {
	var res []int64
	step := int64(1)
	a := id.First
	b := id.Last
	if a > b {
		return res
	}
	for i := a; i != b+step; i += step {
		res = append(res, i)
	}
	return res
}

func (id Pair) SumInvalid() int64 {
	r := int64(0)
	for _, n := range id.invalid {
		r += n
	}
	return r
}

func NewIdPair() Pair {
	return Pair{
		First:   0,
		Last:    0,
		invalid: []int64{},
	}
}

func validInt64(n int64) bool {
	s := strconv.FormatInt(n, 10)
	return valid(s, 2, true) //part 2: at-least 2 repeats
}

func valid(s string, minRepeats int, atLeast bool) bool {
	n := len(s)
	if n < 2 || minRepeats < 2 {
		return true
	}

	for d := 1; d*d <= n; d++ {
		if n%d != 0 {
			continue
		}

		for _, blockLen := range []int{d, n / d} {
			if blockLen >= n {
				continue
			}
			reps := n / blockLen

			// exact count or at-least??
			if !atLeast && reps != minRepeats {
				continue
			}
			if atLeast && reps < minRepeats {
				continue
			}

			block := s[:blockLen]
			match := true
			for i := blockLen; i < n; i += blockLen {
				if s[i:i+blockLen] != block {
					match = false
					break
				}
			}
			if match {
				return false // nope, invalid
			}
		}
	}
	return true
}
