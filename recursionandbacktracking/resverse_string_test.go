package recursionandbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "elppa", ReverseString("apple"))
	assert.Equal(t, "", ReverseString(""))
	assert.Equal(t, "b", ReverseString("b"))
	assert.Equal(t, "sumatopoppih", ReverseString("hippopotamus"))
}
