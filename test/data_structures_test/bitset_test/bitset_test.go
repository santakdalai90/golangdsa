package bitset_test

import (
	"github.com/santakdalai90/golang-dsa/src/data_structures/bitset"
	"testing"

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
