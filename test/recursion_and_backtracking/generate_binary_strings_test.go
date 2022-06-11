package recursion_and_backtracking

import (
	//"github.com/stretchr/testify/assert"
	"golang-dsa/src/recursion_and_backtracking"
	"testing"
)

func TestBinaryString(t *testing.T) {
	t.Log("Binary strings of length 5:", recursion_and_backtracking.BinaryString(5))
	t.Log("Binary strings of length 3:", recursion_and_backtracking.BinaryString(3))
}
