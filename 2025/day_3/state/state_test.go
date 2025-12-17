package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCalculate(t *testing.T) {
	sut := NewState(3)
	sut.topK = []int{3, 5, 2}
	result := sut.Calculate()
	assert.Equal(t, 352, result, "should be equal")
}

func TestShouldReset(t *testing.T) {
	sut := NewState(2)
	sut.topK = []int{3, 5}
	sut.Reset()

	assert.Equal(t, []int{0, 0}, sut.topK)
}

func TestShouldCalculateTopK(t *testing.T) {
	sut := NewState(2)

	arr := []int{8, 1, 9}

	sut.SetInputArray(arr)
	sut.TopK()

	assert.Equal(t, []int{8, 9}, sut.topK)
}

func TestShouldCalculateTopK2(t *testing.T) {
	sut := NewState(2)

	arr := []int{2, 3, 4, 7, 8}

	sut.SetInputArray(arr)
	sut.TopK()

	assert.Equal(t, []int{7, 8}, sut.topK)
}

func TestShouldCalculateTopK3(t *testing.T) {
	sut := NewState(2)

	arr := []int{8, 1, 9, 2}

	sut.SetInputArray(arr)
	sut.TopK()

	assert.Equal(t, []int{9, 2}, sut.topK)
}

func TestShouldCalculateTopK4(t *testing.T) {
	sut := NewState(2)
	arr := []int{3, 7, 3, 3, 4, 4, 4, 4, 4, 4, 3, 3, 7, 2, 4, 4, 3, 4, 1, 4, 6, 3, 4, 5, 2, 4, 6, 3, 6, 4, 4, 2, 3, 4, 4, 9, 3, 3, 5, 4, 1, 4, 4, 5, 8, 4, 4, 3, 3, 3, 4, 4, 4, 2, 5, 4, 4, 4, 4, 5, 3, 5, 3, 4, 4, 5, 4, 4, 4, 4, 3, 4, 3, 4, 4, 4, 3, 2, 4, 3, 3, 5, 4, 5, 4, 4, 4, 6, 4, 2, 3, 3, 4, 3, 4, 4, 4, 4, 7, 2}
	sut.SetInputArray(arr)
	sut.TopK()

	assert.Equal(t, []int{9, 8}, sut.topK)
}

func TestShouldCalculateTopK5(t *testing.T) {
	sut := NewState(2)
	arr := []int{2, 1, 3, 2, 2, 2, 2, 3, 4, 1, 2, 2, 4, 2, 2, 2, 4, 1, 1, 1, 2, 5, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 3, 2, 2, 2, 2, 2, 1, 2, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 2, 2, 2, 2, 1, 2, 2, 2, 6, 2, 2, 1, 2, 2, 2, 1, 4, 2, 2, 1, 2, 2, 2, 2, 1, 2, 2, 3, 2, 2, 1, 4, 2, 1, 1, 1, 2, 2, 5, 1, 4}
	sut.SetInputArray(arr)
	sut.TopK()
	assert.Equal(t, []int{6, 5}, sut.topK)

}

func TestShouldCalculateTopK6(t *testing.T) {
	sut := NewState(2)
	arr := []int{2, 2, 3, 2, 2, 2, 4, 3, 3, 2, 3, 2, 2, 2, 4, 5, 3, 1, 2, 2, 3, 3, 2, 1, 3, 2, 3, 3, 5, 3, 2, 1, 2, 4, 4, 3, 2, 2, 3, 4, 4, 4, 2, 2, 3, 3, 1, 2, 3, 2, 2, 2, 4, 2, 3, 2, 5, 2, 3, 3, 2, 2, 4, 2, 2, 3, 3, 3, 3, 3, 3, 2, 2, 2, 3, 4, 3, 4, 3, 3, 4, 2, 3, 2, 2, 1, 1, 3, 2, 3, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3}
	sut.SetInputArray(arr)
	sut.TopK()
	assert.Equal(t, []int{5, 5}, sut.topK)

}

func TestShouldCalculateTopK7(t *testing.T) {
	sut := NewState(2)
	arr := []int{3, 2, 3, 3, 3, 3, 2, 3, 2, 3, 3, 3, 2, 3, 3, 1, 1, 4, 2, 3, 3, 3, 3, 2, 3, 5, 2, 3, 3, 2, 3, 2, 2, 3, 2, 2, 3, 2, 3, 4, 3, 2, 3, 3, 3, 5, 3, 3, 3, 5, 3, 3, 2, 2, 5, 2, 3, 7, 3, 3, 3, 3, 3, 3, 4, 3, 2, 3, 3, 2, 6, 3, 2, 3, 3, 2, 2, 3, 3, 2, 3, 1, 3, 3, 1, 3, 3, 2, 3, 3}
	sut.SetInputArray(arr)
	sut.TopK()
	assert.Equal(t, []int{7, 6}, sut.topK)

}
