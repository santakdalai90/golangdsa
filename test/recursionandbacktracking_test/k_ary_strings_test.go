package recursionandbacktracking_test

import (
	//"github.com/stretchr/testify/assert"
	"testing"

	"github.com/santakdalai90/golangdsa/src/recursionandbacktracking"
)

func TestKAryString(t *testing.T) {
	t.Log("K-Ary strings of length 3:", recursionandbacktracking.KAryString([]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}, 3))
	t.Log("K-Ary strings of length 5:", recursionandbacktracking.KAryString([]byte{
		'a', 'b', 'c',
	}, 5))
}
