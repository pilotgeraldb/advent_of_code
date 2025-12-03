package id

import (
	"strconv"
	"strings"
)

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
		//nope
		return true
	}

	for d := 1; d*d <= n; d++ { // every possible divisor of the string length
		//consider string length 5, and the following d values:
		//1*1=1; 1 <= 5; run loop
		//2*2=4; 4 <= 5; run loop
		//3*3=9; 9 !<= 5; break
		//d only goes up to the whole number part of the sqrt(n) = 2.2360679775

		// if d does not go into n evenly skip it
		// this check also ensures only whole number block lengths are used
		// if we try to use 2 (from above) here, then 5/2 = 2.5 which is not a whole number
		if n%d != 0 {
			// d (iterator) = 2 and n (string length) = 5 then  []int{1, 2/5}
			continue // d does not go into n evenly
		}

		// because of check above this loop only runs on whole numbers
		// example: d (iterator) = 1 and n (string length) = 5 then  []int{1, 5}
		for _, blockLen := range []int{d, n / d} {

			// if d = 1 then on the first iteration here, blockLen = 1 and n = 5
			// we skip single character blocks
			// since "is the whole string repeated 1 time?" is always true,
			// match would be true, and the function would return false
			// every string is 1 copy of itself so that would make all strings invalid
			if blockLen >= n {
				continue //single block
			}

			// reps are how many times a block repeats to fill the entire string
			// because of check above this loop only runs on whole numbers so we get the number of full blocks that fit
			// example: d (iterator) = 1 and n (string length) = 5 and blockLen = 5 then reps = 5 / 5 = 1
			reps := n / blockLen

			// check min repeat setting exact/at-least
			if !atLeast && reps != minRepeats {
				continue // exact
			}
			if atLeast && reps < minRepeats {
				continue // at-least
			}

			// we define a "block" as a contiguous substring of length `blockLen`
			// when a block is repeated exactly (n / blockLen) times it reconstructs the entire string.
			block := s[:blockLen]

			// ensure the rest of the string is created by repeating the same block
			// iterate the string by blockLen and compare each group
			// even though strings.Repeat(...) is allocation-free and fast,
			// still O(n) for every divisor, worst O(n √n), big oof on strings (over 1 million digits)
			if strings.Repeat(block, reps) == s {
				return false
			}
		}
	}
	return true
}
