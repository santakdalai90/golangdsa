package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testArr := []int{2, 14, 17, 20, 42, 50}
	BinarySearchIterative = true

	idx := BinarySearch(testArr, 15)
	assert.Equal(t, -1, idx)

	idx = BinarySearch(testArr, 20)
	assert.Equal(t, 3, idx)

	BinarySearchIterative = false

	idx = BinarySearch(testArr, 15)
	assert.Equal(t, -1, idx)

	idx = BinarySearch(testArr, 17)
	assert.Equal(t, 2, idx)

	idx = BinarySearch(testArr, 2)
	assert.Equal(t, 0, idx)

	idx = BinarySearch(testArr, 14)
	assert.Equal(t, 1, idx)

	idx = BinarySearch(testArr, 20)
	assert.Equal(t, 3, idx)

	idx = BinarySearch(testArr, 42)
	assert.Equal(t, 4, idx)

	idx = BinarySearch(testArr, 50)
	assert.Equal(t, 5, idx)
}
