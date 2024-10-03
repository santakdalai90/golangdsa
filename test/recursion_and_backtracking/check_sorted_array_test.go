package recursion_and_backtracking

import (
	"testing"

	"github.com/santakdalai90/GolangDSA/src/recursion_and_backtracking"
	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	assert.Equal(t, false, recursion_and_backtracking.IsSorted([]int{4, 2, 1, 7, 3}))
	assert.Equal(t, true, recursion_and_backtracking.IsSorted([]int{1, 2, 3, 4, 7}))
}
