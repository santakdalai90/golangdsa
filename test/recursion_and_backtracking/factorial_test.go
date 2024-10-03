package recursion_and_backtracking

import (
	"github.com/stretchr/testify/assert"
	"github.com/santakdalai90/golang-dsa/src/recursion_and_backtracking"
	"testing"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, recursion_and_backtracking.Factorial(5))
    assert.Equal(t, 720, recursion_and_backtracking.Factorial(6))
    assert.Equal(t, 120, recursion_and_backtracking.FactorialIterative(5))
    assert.Equal(t, 720, recursion_and_backtracking.FactorialIterative(6))
    
}