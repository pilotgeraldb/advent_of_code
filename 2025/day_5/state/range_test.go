package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFindUniqueCount(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(12, 18, true)) // 7
	sut.AddRange(NewRange(10, 14, true)) // 5 -> 2 (because 12, 13, and 14 overlap)

	//total should be 3 + 7 + 2 = 12

	actual := sut.UniqueCount()

	assert.Equal(t, 12, actual, "should be equal")
}

func TestShouldFindUniqueCount2(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true)) // 3
	sut.AddRange(NewRange(3, 5, true)) // 0

	//total should be 3 + 0 = 3

	actual := sut.UniqueCount()

	assert.Equal(t, 3, actual, "should be equal")
}

func TestShouldFindUniqueCount3(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true)) // 3
	sut.AddRange(NewRange(3, 5, true)) // 0
	sut.AddRange(NewRange(6, 8, true)) // 3

	//total should be 3 + 0 + 3 = 6

	actual := sut.UniqueCount()

	assert.Equal(t, 6, actual, "should be equal")
}

func TestShouldFindUniqueCount4(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 3, true)) // 1
	sut.AddRange(NewRange(3, 5, true)) // 2

	//total should be 1 + 2 = 3

	actual := sut.UniqueCount()

	assert.Equal(t, 3, actual, "should be equal")
}

func TestShouldFindUniqueCount5(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true)) // 3
	sut.AddRange(NewRange(3, 5, true)) // 0
	sut.AddRange(NewRange(4, 5, true)) // 0

	//total should be 3 + 0 + 0 = 3

	actual := sut.UniqueCount()

	assert.Equal(t, 3, actual, "should be equal")
}

func TestShouldFindUniqueCount6(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true)) // 3
	sut.AddRange(NewRange(3, 5, true)) // 0
	sut.AddRange(NewRange(4, 6, true)) // 1

	//total should be 3 + 0 + 1 = 4

	actual := sut.UniqueCount()

	assert.Equal(t, 4, actual, "should be equal")
}

func TestShouldFindUniqueCount7(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(8, 11, true))  // 4
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7
	sut.AddRange(NewRange(1, 4, true))   // 2

	actual := sut.UniqueCount()

	assert.Equal(t, 18, actual, "should be equal")
}

func TestShouldFindUniqueCount8(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7

	actual := sut.UniqueCount()

	assert.Equal(t, 14, actual, "should be equal")
}

func TestShouldFindUniqueCount9(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 3, true)) // 1

	actual := sut.UniqueCount()

	assert.Equal(t, 1, actual, "should be equal")
}

func TestShouldFindUniqueCount10(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7

	actual := sut.UniqueCount()

	assert.Equal(t, 14, actual, "should be equal")
}

func TestShouldFindUniqueCount11(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0

	actual := sut.UniqueCount()

	assert.Equal(t, 14, actual, "should be equal")
}

func TestShouldFindUniqueCount12(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(12, 18, true)) // 7
	sut.AddRange(NewRange(10, 14, true)) // 5
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(3, 5, true))   // 3
	sut.AddRange(NewRange(10, 14, true)) // 0
	sut.AddRange(NewRange(16, 20, true)) // 5
	sut.AddRange(NewRange(10, 14, true)) // 0

	actual := sut.UniqueCount()

	assert.Equal(t, 14, actual, "should be equal")
}

func TestShouldFindUniqueCount13(t *testing.T) {
	sut := NewState()

	sut.AddRange(NewRange(10610054450932, 16248915699091, true))
	sut.AddRange(NewRange(377085883505177, 377411515098475, true))
	sut.AddRange(NewRange(375819323667708, 376258116861038, true))
	sut.AddRange(NewRange(170187498282307, 170497327253346, true))
	sut.AddRange(NewRange(166652278386204, 167578065989335, true))
	sut.AddRange(NewRange(553257284780264, 558329119300938, true))
	sut.AddRange(NewRange(42434578175138, 42434578175138, true))

	actual := sut.UniqueCount()

	assert.Equal(t, 12710737129638, actual, "should be equal")
}
