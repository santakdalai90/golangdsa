package recursionandbacktracking_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "elppa", recursionandbacktracking.ReverseString("apple"))
	assert.Equal(t, "", recursionandbacktracking.ReverseString(""))
	assert.Equal(t, "b", recursionandbacktracking.ReverseString("b"))
	assert.Equal(t, "sumatopoppih", recursionandbacktracking.ReverseString("hippopotamus"))
}
