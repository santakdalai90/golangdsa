package recursion_and_backtracking

import (
	"testing"

	"github.com/santakdalai90/GolangDSA/src/recursion_and_backtracking"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "elppa", recursion_and_backtracking.ReverseString("apple"))
	assert.Equal(t, "", recursion_and_backtracking.ReverseString(""))
	assert.Equal(t, "b", recursion_and_backtracking.ReverseString("b"))
	assert.Equal(t, "sumatopoppih", recursion_and_backtracking.ReverseString("hippopotamus"))
}
