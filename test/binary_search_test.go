package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/santakdalai90/golang-dsa/src"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testArr := []int{2, 14, 17, 20, 42, 50}
	src.BinarySearchIterative = true

	idx := src.BinarySearch(testArr, 15)
	assert.Equal(t, -1, idx)

	idx = src.BinarySearch(testArr, 20)
	assert.Equal(t, 3, idx)

	src.BinarySearchIterative = false

	idx = src.BinarySearch(testArr, 15)
	assert.Equal(t, -1, idx)

	idx = src.BinarySearch(testArr, 17)
	assert.Equal(t, 2, idx)

	idx = src.BinarySearch(testArr, 2)
	assert.Equal(t, 0, idx)

	idx = src.BinarySearch(testArr, 14)
	assert.Equal(t, 1, idx)

	idx = src.BinarySearch(testArr, 20)
	assert.Equal(t, 3, idx)

	idx = src.BinarySearch(testArr, 42)
	assert.Equal(t, 4, idx)

	idx = src.BinarySearch(testArr, 50)
	assert.Equal(t, 5, idx)
}
