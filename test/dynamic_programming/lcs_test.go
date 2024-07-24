package dynamicprogramming

import (
	dynamicprogramming "golang-dsa/src/dynamic_programming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCS(t *testing.T) {
	lcsLen, lcs := dynamicprogramming.LCS(
		"abcabba",
		"cbabaca",
	)
	assert.Equal(t, 4, lcsLen)
	assert.Equal(t, "cbba", lcs)

	lcsLen, lcs = dynamicprogramming.LCS(
		"ataatctgtactttcgcgcgcgatacagttgctcaagtagccattcgggtgcctgacggcatgacatgtatgccactttcaccaaagtaatcatcttaac",
		"ggtaataaaaggctcgcccaccccgaccctgaagtgaagaagtaacgttgttcatgatcgacccgcatgaattcatagcaacgccgactgatctactccc",
	)
	assert.Equal(t, 62, lcsLen)
	assert.Equal(t, "taatactcgcccgataagtgaagaatcgtgctgacggcatgaattatgcaaccgacatctac", lcs)
}

func TestLCSOptimized(t *testing.T) {
	lcsLen := dynamicprogramming.LCSOptimized(
		"abcabba",
		"cbabaca",
	)
	assert.Equal(t, 4, lcsLen)

	lcsLen = dynamicprogramming.LCSOptimized(
		"ataatctgtactttcgcgcgcgatacagttgctcaagtagccattcgggtgcctgacggcatgacatgtatgccactttcaccaaagtaatcatcttaac",
		"ggtaataaaaggctcgcccaccccgaccctgaagtgaagaagtaacgttgttcatgatcgacccgcatgaattcatagcaacgccgactgatctactccc",
	)
	assert.Equal(t, 62, lcsLen)
}
