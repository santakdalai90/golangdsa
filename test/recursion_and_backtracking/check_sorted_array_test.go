package recursion_and_backtracking

import (
	"github.com/stretchr/testify/assert"
	"golang-dsa/src/recursion_and_backtracking"
	"testing"
)

func TestIsSorted(t *testing.T) {
	assert.Equal(t, false, recursion_and_backtracking.IsSorted([]int{4, 2, 1, 7, 3}))
	assert.Equal(t, true, recursion_and_backtracking.IsSorted([]int{1, 2, 3, 4, 7}))
}
