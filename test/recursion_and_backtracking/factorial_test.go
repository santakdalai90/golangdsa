package recursion_and_backtracking

import (
	"testing"

	"github.com/santakdalai90/GolangDSA/src/recursion_and_backtracking"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, recursion_and_backtracking.Factorial(5))
	assert.Equal(t, 720, recursion_and_backtracking.Factorial(6))
	assert.Equal(t, 120, recursion_and_backtracking.FactorialIterative(5))
	assert.Equal(t, 720, recursion_and_backtracking.FactorialIterative(6))

}
