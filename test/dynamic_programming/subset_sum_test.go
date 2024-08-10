package dynamicprogramming

import (
	dynamicprogramming "golang-dsa/src/dynamic_programming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetSum(t *testing.T) {
	tests := []struct {
		name   string
		set    []int
		sum    int
		exists bool
	}{
		{"simple", []int{3, 34, 4, 12, 5, 2}, 9, true},
		{"doesn't exsit", []int{3, 34, 4, 12, 5, 2}, 30, false},
		{"empty array", []int{}, 0, true},
		{"single element exists", []int{10}, 10, true},
		{"single element doesn't exist", []int{10}, 5, false},
		{"large array exists", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55, true},
		{"large array doesn't exist", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 60, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exists := dynamicprogramming.IsSubsetSum(test.set, test.sum)
			assert.Equal(t, test.exists, exists)
		})
	}
}
