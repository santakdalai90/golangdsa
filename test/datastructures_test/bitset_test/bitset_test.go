package bitset_test

import (
	"testing"

	"github.com/santakdalai90/golangdsa/src/datastructures/bitset"

	"github.com/stretchr/testify/assert"
)

func TestBitSet(t *testing.T) {
	bs := bitset.NewBitSet(128)

	bs.Set(5)
	bs.Set(100)

	assert.Equal(t, true, bs.IsSet(5))
	assert.Equal(t, true, bs.IsSet(100))
	assert.Equal(t, false, bs.IsSet(64))

	bs.Clear(5)
	assert.Equal(t, false, bs.IsSet(5))
}
