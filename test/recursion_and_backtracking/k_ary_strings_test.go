package recursion_and_backtracking

import (
	//"github.com/stretchr/testify/assert"
	"github.com/santakdalai90/golang-dsa/src/recursion_and_backtracking"
	"testing"
)

func TestKAryString(t *testing.T) {
	t.Log("K-Ary strings of length 3:", recursion_and_backtracking.KAryString([]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}, 3))
	t.Log("K-Ary strings of length 5:", recursion_and_backtracking.KAryString([]byte{
		'a', 'b', 'c',
	}, 5))
}
