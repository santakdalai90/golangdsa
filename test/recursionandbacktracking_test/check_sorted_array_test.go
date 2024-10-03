package recursionandbacktracking_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	assert.Equal(t, false, recursionandbacktracking.IsSorted([]int{4, 2, 1, 7, 3}))
	assert.Equal(t, true, recursionandbacktracking.IsSorted([]int{1, 2, 3, 4, 7}))
}
