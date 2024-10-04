package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCS(t *testing.T) {
	lcsLen, lcs := LCS(
		"abcabba",
		"cbabaca",
	)
	assert.Equal(t, 4, lcsLen)
	assert.Equal(t, "cbba", lcs)

	lcsLen, lcs = LCS(
		"ataatctgtactttcgcgcgcgatacagttgctcaagtagccattcgggtgcctgacggcatgacatgtatgccactttcaccaaagtaatcatcttaac",
		"ggtaataaaaggctcgcccaccccgaccctgaagtgaagaagtaacgttgttcatgatcgacccgcatgaattcatagcaacgccgactgatctactccc",
	)
	assert.Equal(t, 62, lcsLen)
	assert.Equal(t, "taatactcgcccgataagtgaagaatcgtgctgacggcatgaattatgcaaccgacatctac", lcs)

	lcsLen, lcs = LCS(
		"abacc",
		"babcacb",
	)
	assert.Equal(t, 4, lcsLen)
	assert.Equal(t, "bacc", lcs)
}

func TestLCSOptimized(t *testing.T) {
	lcsLen := LCSOptimized(
		"abcabba",
		"cbabaca",
	)
	assert.Equal(t, 4, lcsLen)

	lcsLen = LCSOptimized(
		"ataatctgtactttcgcgcgcgatacagttgctcaagtagccattcgggtgcctgacggcatgacatgtatgccactttcaccaaagtaatcatcttaac",
		"ggtaataaaaggctcgcccaccccgaccctgaagtgaagaagtaacgttgttcatgatcgacccgcatgaattcatagcaacgccgactgatctactccc",
	)
	assert.Equal(t, 62, lcsLen)
}
