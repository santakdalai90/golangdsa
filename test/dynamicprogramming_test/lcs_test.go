package dynamicprogramming_test

import (
	"testing"

	dynamicprogramming "github.com/santakdalai90/golangdsa/src/dynamicprogramming"

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

	lcsLen, lcs = dynamicprogramming.LCS(
		"abacc",
		"babcacb",
	)
	assert.Equal(t, 4, lcsLen)
	assert.Equal(t, "bacc", lcs)
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
