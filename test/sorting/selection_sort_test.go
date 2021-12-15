package sorting

import (
	"github.com/stretchr/testify/assert"
	"golang-dsa/src/sorting"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{4, 2, 1, 7, 3}
	sorting.SelectionSort(a)
	assert.Equal(t, []int{1, 2, 3, 4, 7}, a)

	a = []int{65, 55, 45, 35, 25, 15, 10}
	sorting.SelectionSort(a)
	assert.Equal(t, []int{10, 15, 25, 35, 45, 55, 65}, a)
}
