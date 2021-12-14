package test

import (
	"github.com/stretchr/testify/assert"
	"golang-dsa/src"
	"testing"
)

func TestArrayMax(t *testing.T) {
	testArr := []int{14, 53, 78, 31, 97, 65, 23, 124, 14, 39, 45}
	max, err := src.ArrayMax(testArr)
	assert.NoError(t, err)
	assert.Equal(t, 124, max)
}

func TestArrayMaxEmptyArray(t *testing.T) {
	testArr := []int{}
	_, err := src.ArrayMax(testArr)
	assert.EqualError(t, err, src.ErrorEmptyArray)
}
