package recursion_and_backtracking

import (
	//"github.com/stretchr/testify/assert"
	"testing"

	"github.com/santakdalai90/GolangDSA/src/recursion_and_backtracking"
)

func TestBinaryString(t *testing.T) {
	t.Log("Binary strings of length 5:", recursion_and_backtracking.BinaryString(5))
	t.Log("Binary strings of length 3:", recursion_and_backtracking.BinaryString(3))
}
