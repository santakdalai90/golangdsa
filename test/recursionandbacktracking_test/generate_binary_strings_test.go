package recursionandbacktracking_test

import (
	//"github.com/stretchr/testify/assert"
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
)

func TestBinaryString(t *testing.T) {
	t.Log("Binary strings of length 5:", recursionandbacktracking.BinaryString(5))
	t.Log("Binary strings of length 3:", recursionandbacktracking.BinaryString(3))
}
