package recursion_and_backtracking

import (
	"github.com/stretchr/testify/assert"
	"golang-dsa/src/recursion_and_backtracking"
	"testing"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "elppa", recursion_and_backtracking.ReverseString("apple"))
    assert.Equal(t, "", recursion_and_backtracking.ReverseString(""))
    assert.Equal(t, "b", recursion_and_backtracking.ReverseString("b"))
    assert.Equal(t, "sumatopoppih", recursion_and_backtracking.ReverseString("hippopotamus"))
}