package main

import (
	"day_2_aoc_2025/id"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdPairValidate(t *testing.T) {
	p := id.Pair{
		First: 11,
		Last:  22,
	}

	r := p.Validate()

	assert.Equal(t, int64(33), r, "should be equal")
}
