package recursionandbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	assert.Equal(t, false, IsSorted([]int{4, 2, 1, 7, 3}))
	assert.Equal(t, true, IsSorted([]int{1, 2, 3, 4, 7}))
}
