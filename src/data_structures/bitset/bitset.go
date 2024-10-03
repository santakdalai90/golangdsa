package bitset

// BitSet represents a vector of bits.
type BitSet struct {
	bits []uint64
}

// NewBitSet creates a new BitSet of the given size.
func NewBitSet(size int) *BitSet {
	// Each uint64 holds 64 bits, so we divide the size by 64
	return &BitSet{
		bits: make([]uint64, (size+63)/64), // Round up to nearest 64-bit chunk
	}
}

// Set sets the bit at the given index to 1.
func (bs *BitSet) Set(index int) {
	// Find the index in the uint64 slice and the bit position
	word, bit := index/64, uint(index%64)
	bs.bits[word] |= 1 << bit
}

// Clear clears the bit at the given index (sets to 0).
func (bs *BitSet) Clear(index int) {
	word, bit := index/64, uint(index%64)
	bs.bits[word] &^= 1 << bit
}

// IsSet checks if the bit at the given index is set to 1.
func (bs *BitSet) IsSet(index int) bool {
	word, bit := index/64, uint(index%64)
	return (bs.bits[word] & (1 << bit)) != 0
}

// Size returns the size of the BitSet.
func (bs *BitSet) Size() int {
	return len(bs.bits) * 64
}
