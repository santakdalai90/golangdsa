package test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "golang-dsa/src"
)

func TestBinarySearch(t *testing.T) {
    testArr := []int{2, 14, 17, 20, 42, 50}

    idx := src.BinarySearch(testArr, 15)
    assert.Equal(t, idx, -1)

    idx = src.BinarySearch(testArr, 20)
    assert.Equal(t, idx, 3)
}